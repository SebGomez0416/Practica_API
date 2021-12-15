package handler

import "net/http"

func RoutePerson(mux *http.ServeMux, s Storage) {

	h := newPerson(s)

	mux.HandleFunc("/v1/persons/create", h.create)
}
