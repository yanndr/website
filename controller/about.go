package controller

import (
	"log"
	"net/http"
)

type about struct {
	controller
}

func (a about) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/about", a.handleAbout)
}

func (a about) handleAbout(w http.ResponseWriter, r *http.Request) {
	err := a.template.Execute(w, nil)
	if err != nil {
		log.Println("error ", err)
	}

}
