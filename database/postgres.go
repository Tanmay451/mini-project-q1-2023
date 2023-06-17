package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() (*sql.DB, error) {
	// Create the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "your-password", "forum")

	// Connect to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
