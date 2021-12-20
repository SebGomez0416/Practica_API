package handler

import "github.com/SebGomez0416/Practica_API/model"

type Storage interface {
	Create(*model.Person) error
	Update(*model.Person) error
	GetAll() (model.Persons, error)
	GetByID(uint) (*model.Person, error)
	Delete(uint) error
}

type StorageLogin interface {
	GetByEmail(string) (*model.Login, error)
	IsLoginValid(data *model.Login) bool
}
