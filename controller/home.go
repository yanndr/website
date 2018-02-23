package controller

import (
	"log"
	"net/http"
)

type home struct {
	controller
}

func (h home) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/home", h.handleHome)
	mux.HandleFunc("/", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	err := h.template.Execute(w, nil)
	if err != nil {
		log.Println("error ", err)
	}

}
