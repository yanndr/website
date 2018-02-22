package controller

import (
	"html/template"
	"log"
	"net/http"
)

type home struct {
	homeTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	err := h.homeTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error ", err)
	}

}
