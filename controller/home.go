package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/yanndr/website/viewmodel"
)

type home struct {
	controller
}

func (h home) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/home", h.handleHome)
	mux.HandleFunc("/", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {

	vm := &viewmodel.Home{YearOfXp: time.Now().Year() - 2001}

	err := h.template.Execute(w, vm)

	if err != nil {
		log.Println("error ", err)
	}

}
