package storage

import (
	"database/sql"
	"fmt"

	"github.com/SebGomez0416/Practica_API/model"
)

const (
	psqlCreatePerson = `INSERT INTO person(name,age,community_name)
	VALUES($1,$2,$3) RETURNING id`

	psqlUpdatePerson = `UPDATE person SET name = $1, age = $2,
	community_name = $3 WHERE id = $4`

	psqlDeletePerson = `DELETE FROM person WHERE id = $1`

	psqlGetAllPerson = `SELECT  id, name, age,community_name
	FROM person`
)

// PsqlPerson used for work whit postgres - person
type PsqlPerson struct {
	db *sql.DB
}

// NewPslProduct return a new pointer of PsqlPerson
func NewPslPerson(db *sql.DB) *PsqlPerson {
	return &PsqlPerson{db}
}

// create implement the interface person.model
func (psql *PsqlPerson) Create(p *model.Person) error {
	stmt, err := psql.db.Prepare(psqlCreatePerson)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(p.Name, p.Age, p.CommunityName).Scan(&p.ID)
	if err != nil {
		return err
	}
	fmt.Println(" se creo la persona correctamente")
	return nil

}

// Update implement the interface person.model
func (psql *PsqlPerson) Update(p *model.Person) error {

	stmt, err := psql.db.Prepare(psqlUpdatePerson)
	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(p.Name, p.Age, p.CommunityName, p.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no existe la persona con el  id: %d", p.ID)
	}
	fmt.Println("se actualizo correctamente")
	return nil
}

// GetAll implement the interface person.model
func (psql *PsqlPerson) GetAll() (model.Persons, error) {

	stmt, err := psql.db.Prepare(psqlGetAllPerson)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ps := make(model.Persons, 0)

	for rows.Next() {
		p := &model.Person{}

		err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.CommunityName)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ps, nil
}

// Delete implement the interface person.model
func (psql *PsqlPerson) Delete(id uint) error {
	stmt, err := psql.db.Prepare(psqlDeletePerson)

	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no existe la persona con id: %d", id)
	}
	fmt.Println("se borro correctamente")
	return nil

}
