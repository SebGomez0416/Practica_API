package middleware

import (
	"net/http"

	"github.com/SebGomez0416/Practica_API/authorization"
)

func Authentication(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)

		if err != nil {
			forbidden(w, r)
			return
		}
		f(w, r)
	}
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("token invalido"))

}
