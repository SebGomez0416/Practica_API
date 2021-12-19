package main

import (
	"log"
	"net/http"

	"github.com/SebGomez0416/Practica_API/authorization"
	"github.com/SebGomez0416/Practica_API/handler"
	"github.com/SebGomez0416/Practica_API/storage"
)

func main() {

	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")

	if err != nil {
		log.Fatalf("no se pudo cargar los certificados: %v", err)
	}

	storage.NewPostgresDB()
	storagePerson := storage.NewPslPerson(storage.Pool())

	mux := http.NewServeMux()

	handler.RoutePerson(mux, storagePerson)
	log.Println("servidor iniciado en el puerto :8080")
	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Printf("error en el servidor: %v\n", err)
	}

}
