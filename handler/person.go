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
		resp := newResponse(Error, "Metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	data := model.Person{}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		resp := newResponse(Error, "la persona no fue creada correctamente", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	err = p.storage.Create(&data)

	if err != nil {
		resp := newResponse(Error, "hubo un problema al crear la persona", nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := newResponse(Message, "persona creada correctamente", nil)
	responseJSON(w, http.StatusCreated, resp)
}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		resp := newResponse(Error, "Metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	data, err := p.storage.GetAll()

	if err != nil {
		resp := newResponse(Error, "hubo un problema al obtener todas las personas", nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := newResponse(Message, "Ok", data)
	responseJSON(w, http.StatusOK, resp)

}

func (p *person) getByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		resp := newResponse(Error, "Metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		resp := newResponse(Error, "error al cargar el id", nil)
		responseJSON(w, http.StatusBadRequest, resp)
	}

	res, err := p.storage.GetByID(data.ID)

	if err != nil {
		resp := newResponse(Error, "hubo un problema al obtener la personas", nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := newResponse(Message, "Ok", res)
	responseJSON(w, http.StatusOK, resp)

}

func (p *person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		resp := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	data := model.Person{}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		resp := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	err = p.storage.Update(&data)

	if err != nil {
		resp := newResponse(Error, "Hubo un problema al obtener todas las personas", nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := newResponse(Message, "Persona actualizada correctamente", nil)
	responseJSON(w, http.StatusOK, resp)

}

func (p *person) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		resp := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	data := model.Person{}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		resp := newResponse(Error, "error al cargar la persona", nil)
		responseJSON(w, http.StatusBadRequest, resp)
	}

	err = p.storage.Delete(data.ID)

	if err != nil {
		resp := newResponse(Error, "Ocurrió un error al elminar el registro", nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := newResponse(Message, "Ok", nil)
	responseJSON(w, http.StatusOK, resp)

}
