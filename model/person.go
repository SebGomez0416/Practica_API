package model

import (
	"fmt"
)

//  person estructura de una persona
type Person struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Age           uint8  `json:"age"`
	CommunityName string `json:"community_name"`
}

// Persons slice de personas
type Persons []*Person

func (p *Person) String() string {
	return fmt.Sprintf("id: %d, name:%s, age: %d, community: %s,\n", p.ID, p.Name, p.Age, p.CommunityName)
}
