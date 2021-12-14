package main

import (
	"fmt"
	"log"

	"github.com/SebGomez0416/Practica_API/model"
	"github.com/SebGomez0416/Practica_API/storage"
)

func main() {

	storage.NewPostgresDB()

	storagePerson := storage.NewPslPerson(storage.Pool())
	servicePerson := model.NewService(storagePerson)

	// p := &model.Person{
	// 	Name:          "carlos santana",
	// 	Age:           23,
	// 	CommunityName: "platzi",
	// }

	ps, err := servicePerson.GetByID(3)

	if err != nil {

		log.Fatalf("Person.Create: %v", err)
	}

	fmt.Println(ps)

}
