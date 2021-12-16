package handler

import "net/http"

func RoutePerson(mux *http.ServeMux, s Storage) {

	h := newPerson(s)

	mux.HandleFunc("/v1/persons/create", h.create)
	mux.HandleFunc("/v1/persons/get-all", h.getAll)
	mux.HandleFunc("/v1/persons/get-by-id", h.getByID)
	mux.HandleFunc("/v1/persons/update", h.update)
	mux.HandleFunc("/v1/persons/delete", h.delete)
}
