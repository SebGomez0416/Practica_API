package storage

import (
	"database/sql"
	"fmt"

	"github.com/SebGomez0416/Practica_API/model"
)

const (
	psqlGetLoginByEmail = `SELECT email, password
	FROM login  WHERE email = $1`
)

type PsqlLogin struct {
	db *sql.DB
}

func NewPsqlLogin(db *sql.DB) *PsqlLogin {
	return &PsqlLogin{db}
}

func (psql *PsqlLogin) GetByEmail(email string) (*model.Login, error) {

	stmt, err := psql.db.Prepare(psqlGetLoginByEmail)
	if err != nil {
		return &model.Login{}, err
	}
	defer stmt.Close()

	return ScanRowLogin(stmt.QueryRow(email))
}

func ScanRowLogin(s scanner) (*model.Login, error) {
	p := &model.Login{}

	err := s.Scan(&p.Email, &p.Password)

	if err != nil {
		return &model.Login{}, err
	}

	return p, err
}

func (psql *PsqlLogin) IsLoginValid(data *model.Login) bool {

	login, err := psql.GetByEmail(data.Email)
	if err != nil {
		fmt.Println("email ingresado incorrecto ")
		return false
	}
	return login.Password == data.Password
}
