package handler

import (
	"encoding/json"
	"net/http"

	"github.com/SebGomez0416/Practica_API/model"
)

type person struct {
	storage Storage
}

func newPerson(s Storage) *person {
	return &person{s}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		w.Header().Set("Content-Type", "application/jason")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type" : "error" , "message" : "Metodo no permitido"`))
		return
	}

	data := model.Person{}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		w.Header().Set("Content-Type", "application/jason")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type" : "error" , "message" : "la persona no fue cargada correctamente"`))
		return
	}

	err = p.storage.Create(&data)

	if err != nil {
		w.Header().Set("Content-Type", "application/jason")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type" : "error" , "message" : "Hubo un problema al crear la persona"`))
		return
	}

	w.Header().Set("Content-Type", "application/jason")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message_type" : "message" , "message" : "persona creada correctamente"`))

}
