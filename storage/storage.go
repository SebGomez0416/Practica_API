package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

//NewPostgresDB  open DataBase
func NewPostgresDB() {
	once.Do(func() {
		var err error

		db, err = sql.Open("postgres", "postgres://postgres:seb95025@localhost:5432/practica_API?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open DB: %v ", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v ", err)
		}
		fmt.Println("conectado a postgres")

	})

}

// pool reutrn a unique instance of db
func Pool() *sql.DB {
	return db
}
