package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Open(dbURL string) (*sql.DB, error) {
	log.Printf("Connecting to database: %s", dbURL)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to database")
	return db, nil
}
