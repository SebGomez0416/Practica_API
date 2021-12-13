package main

import (
	"log"

	"github.com/SebGomez0416/Practica_API/model"
	"github.com/SebGomez0416/Practica_API/storage"
)

func main() {

	storage.NewPostgresDB()

	storagePerson := storage.NewPslPerson(storage.Pool())
	servicePerson := model.NewService(storagePerson)

	p := &model.Person{ID: 1,
		Name:          "sebastian gomez",
		Age:           30,
		CommunityName: "EDTeam",
	}

	err := servicePerson.Update(p)

	if err != nil {

		log.Fatalf("Person.Create: %v", err)
	}

}
