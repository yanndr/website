package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"github.com/yanndr/website/pkg/model"
)

type handler struct {
	templates map[string]*template.Template
	pr        model.ProfileRepository
}

//New return a new handler for the website.
func New(templates map[string]*template.Template, repo model.ProfileRepository) http.Handler {
	h := &handler{templates, repo}

	mux := httprouter.New()
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger(), negroni.NewStatic(http.Dir("public")))
	n.UseHandler(mux)

	mux.GET("/", h.home)
	mux.GET("/home", h.home)
	mux.GET("/resume", h.resume)

	return n
}

func (h *handler) home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	p, err := h.pr.Get()
	if err != nil {
		log.Printf("error on home page: %s", err)
		return
	}
	err = h.templates["home.html"].Execute(w, p)

	if err != nil {
		log.Println("error execting template: ", err)
	}
}

func (h *handler) resume(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	p, err := h.pr.Get()
	if err != nil {
		log.Printf("error on home page: %s", err)
		return
	}
	err = h.templates["resume.html"].Execute(w, p)

	if err != nil {
		log.Println("error ", err)
	}
}
