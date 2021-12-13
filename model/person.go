package model

//  person estructura de una persona
type Person struct {
	ID            int
	Name          string
	Age           uint8
	CommunityName string
}

// Persons slice de personas
type Persons []*Person

type Storage interface {
	Create(*Person) error
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
