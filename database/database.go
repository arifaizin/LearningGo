package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Successfully connected to the database")

	return db, nil
}
