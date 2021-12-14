package model

import (
	"errors"
	"fmt"
)

//  person estructura de una persona
type Person struct {
	ID            uint
	Name          string
	Age           uint8
	CommunityName string
}

// Persons slice de personas
type Persons []*Person

type Storage interface {
	Create(*Person) error
	Update(*Person) error
	GetAll() (Persons, error)
	GetByID(uint) (*Person, error)
	Delete(uint) error
}

func (p *Person) String() string {

	return fmt.Sprintf("id: %d, name:%s, age: %d, community: %s,\n", p.ID, p.Name, p.Age, p.CommunityName)
}

//servicio de product
type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Create(p *Person) error {
	return s.storage.Create(p)
}

func (s *Service) Update(m *Person) error {

	if m.ID == 0 {
		return errors.New("la persona no contien un id")
	}

	return s.storage.Update(m)
}

func (s *Service) GetAll() (Persons, error) {
	return s.storage.GetAll()
}

func (s *Service) GetByID(id uint) (*Person, error) {
	return s.storage.GetByID(id)
}

func (s *Service) Delete(id uint) error {

	return s.storage.Delete(id)
}
