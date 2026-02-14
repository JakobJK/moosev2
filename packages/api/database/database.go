package database

import (
	"database/sql"
	"log"
	"moose-api/db_gen"

	_ "github.com/lib/pq"
)

func Init() (*db_gen.Queries, *sql.DB) {
	connStr := "postgres://moose:moose@db:5432/moose_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db_gen.New(db), db
}
