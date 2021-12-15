package model

import (
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

func (p *Person) String() string {
	return fmt.Sprintf("id: %d, name:%s, age: %d, community: %s,\n", p.ID, p.Name, p.Age, p.CommunityName)
}
