package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/yanndr/website/pkg/handler"
	"github.com/yanndr/website/pkg/model"
	"github.com/yanndr/website/pkg/repository/file"
)

//Version of the program.
var Version = "No Version Provided"

//Build is the GitHash of the Program.
var Build = "No GitHash Provided"

var templates map[string]*template.Template
var mainTmpl = `{{define "main" }} {{ template "base" . }} {{ end }}`

var p *model.Profile

func main() {

	var config struct {
		Port      string `default:"8080"`
		Templates string `default:"templates/"`
		DB        string `default:"profile.json"`
	}

	if err := envconfig.Process("", &config); err != nil {
		log.Print(err)
		//envconfig.Usage("", &config)
		os.Exit(1)
	}

	f, err := os.OpenFile(config.DB, os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open or create file %s: %s", config.DB, err)
		os.Exit(1)
	}

	r := file.NewProfileRepository(f)

	r.Get()

	templateNeedUpdate(config.Templates)
	loadTemplates(config.Templates)
	h := handler.New(templates, r)
	srv := &http.Server{Addr: ":" + config.Port, Handler: h}

	go func() {
		var lastDataFileupdate = time.Unix(0, 0)
		var needUpdate bool
		for range time.Tick(300 * time.Millisecond) {
			isUpdated := templateNeedUpdate(config.Templates)
			if isUpdated {
				log.Println("updating templates")
				loadTemplates(config.Templates)
			}

			if needUpdate, lastDataFileupdate = fileNeedRefresh(config.DB, lastDataFileupdate); needUpdate {
				log.Println("refresh data ")
				p, err = r.Get()
				if err != nil {
					log.Fatalf("could not get profile: %s", err)
				}
			}
		}
	}()

	go func() {
		// graceful shutdown
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		<-interrupt
		log.Print("app is shutting down...")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("could not shutdown: %v\n", err)
		}
	}()

	log.Println("Website version: ", Version, " - ", Build)
	log.Printf("app is ready to listen and serve on port %s", config.Port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("server failed: %v", err)
		os.Exit(1)
	}

	log.Print("good bye!")
}

var lastModTime = time.Unix(0, 0)

func loadTemplates(templateDir string) {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layoutFiles, err := filepath.Glob(templateDir + "shared/" + "*.html")
	if err != nil {
		log.Fatal(err)
	}

	includeFiles, err := filepath.Glob(templateDir + "*.html")
	if err != nil {
		log.Fatal(err)
	}

	fm := template.FuncMap{"type": Type}
	mainTemplate := template.New("main").Funcs(fm)

	mainTemplate, err = mainTemplate.Parse(mainTmpl)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			log.Fatal(err)
		}
		templates[fileName] = template.Must(templates[fileName].ParseFiles(files...))
	}

	log.Println("templates loading successful")

}

func templateNeedUpdate(basePath string) bool {
	needUpdate := false

	f, _ := os.Open(basePath)

	fileInfos, _ := f.Readdir(-1)

	for _, fi := range fileInfos {
		if fi.IsDir() {
			if templateNeedUpdate(basePath + fi.Name() + "/") {
				return true
			}
		}
		if fi.ModTime().After(lastModTime) {
			lastModTime = fi.ModTime()
			needUpdate = true
		}
	}
	return needUpdate
}

func fileNeedRefresh(path string, lastupdate time.Time) (bool, time.Time) {

	fi, err := os.Stat(path)
	if err != nil {
		log.Printf("could not get info for the file %s: %s", path, err)
	}
	return fi.ModTime().After(lastupdate), fi.ModTime()
}

func Type(obj interface{}) string {
	return reflect.TypeOf(obj).String()
}
