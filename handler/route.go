package handler

import (
	"net/http"

	"github.com/SebGomez0416/Practica_API/middleware"
)

func RoutePerson(mux *http.ServeMux, s Storage) {

	h := newPerson(s)

	mux.HandleFunc("/v1/persons/create", middleware.Authentication(h.create))
	mux.HandleFunc("/v1/persons/get-all", middleware.Authentication(h.getAll))
	mux.HandleFunc("/v1/persons/get-by-id", middleware.Authentication(h.getByID))
	mux.HandleFunc("/v1/persons/update", middleware.Authentication(h.update))
	mux.HandleFunc("/v1/persons/delete", middleware.Authentication(h.delete))
}
