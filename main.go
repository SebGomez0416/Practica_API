package main

import (
	"github.com/SebGomez0416/Practica_API/storage"
)

func main() {

	storage.NewPostgresDB()

	// storagePerson := storage.NewPslPerson(storage.Pool())
	// servicePerson := model.NewService(storagePerson)

	// err := servicePerson.Create()

	// if err != nil {

	// 	log.Fatalf("Person.Create: %v", err)
	// }

}
