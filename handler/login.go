package handler

import (
	"encoding/json"
	"net/http"

	"github.com/SebGomez0416/Practica_API/authorization"
	"github.com/SebGomez0416/Practica_API/model"
)

type login struct {
	storage StorageLogin
}

func newLogin(s StorageLogin) *login {
	return &login{s}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		resp := newResponse(Error, "Metodo no permitido", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp := newResponse(Error, "estructura no permitida", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	if !l.storage.IsLoginValid(&data) {
		resp := newResponse(Error, "usuario o contraseña incorrectos", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	token, err := authorization.GenerateToken(&data)

	if err != nil {
		resp := newResponse(Error, "erro al crear el token ", nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	dataToken := map[string]string{"token": token}
	resp := newResponse(Message, "OK", dataToken)
	responseJSON(w, http.StatusOK, resp)

}
