package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := "user=postgres dbname=postgres password=0000 host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Database connection successfully established.")

	return db, nil
}