package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbConnect() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:abdulaziz1221@localhost:5432/n9?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
