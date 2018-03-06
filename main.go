package main

import (
	"context"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	"github.com/urfave/negroni"
	"github.com/yanndr/website/viewmodel"
)

//Version of the program.
var Version = "No Version Provided"

//Build is the GitHash of the Program.
var Build = "No GitHash Provided"

var templates map[string]*template.Template

func main() {

	var config struct {
		Port      string `default:"8080"`
		Templates string `default:"templates"`
	}

	if err := envconfig.Process("", &config); err != nil {
		log.Print(err)
		//envconfig.Usage("", &config)
		os.Exit(1)
	}

	templates = populateTemplates(config.Templates)

	mux := httprouter.New()
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger(), negroni.NewStatic(http.Dir("public")))
	n.UseHandler(mux)

	mux.GET("/", home)
	mux.GET("/home", home)
	mux.GET("/about", about)

	srv := &http.Server{Addr: ":" + config.Port, Handler: n}

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

	go func() {
		for range time.Tick(300 * time.Millisecond) {
			isUpdated := templateNeedUpdate(config.Templates)
			if isUpdated {
				log.Println("updating templates")
				templates = populateTemplates(config.Templates)
			}
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

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	err := templates["home.html"].Execute(w, nil)

	if err != nil {
		log.Println("error ", err)
	}
}

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	vm := &viewmodel.Home{YearOfXp: time.Now().Year() - 2001}

	err := templates["about.html"].Execute(w, vm)

	if err != nil {
		log.Println("error ", err)
	}
}

var lastModTime = time.Unix(0, 0)

func populateTemplates(basePath string) map[string]*template.Template {

	result := make(map[string]*template.Template)
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_nav.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}

func templateNeedUpdate(basePath string) bool {
	needUpdate := false

	f, _ := os.Open(basePath)

	fileInfos, _ := f.Readdir(-1)
	for _, fi := range fileInfos {
		if fi.ModTime().After(lastModTime) {
			lastModTime = fi.ModTime()
			needUpdate = true
		}
	}
	return needUpdate
}
