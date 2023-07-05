package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func CreateDatabaseConnection() (*sql.DB, error) {
	connectionString := "user=postgres dbname=postgres password=0000 host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connection successfully established")

	return db, nil
}

func CreateTestDatabaseConnection() (*sql.DB, error) {
	connectionString := "user=postgres dbname=tests password=0000 host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Test database connection successfully established")

	return db, nil
}
