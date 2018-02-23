package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController  home
	aboutController about
)

type controller struct {
	template *template.Template
}

type routeRegisterer interface {
	registerRoutes()
}

var routes []routeRegisterer

func Startup(mux *http.ServeMux, templates map[string]*template.Template) {

	homeController.template = templates["home.html"]
	homeController.registerRoutes(mux)
	aboutController.template = templates["about.html"]
	aboutController.registerRoutes(mux)
	// http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("wwwroot")))
	http.Handle("/js/", http.FileServer(http.Dir("wwwroot")))
}
