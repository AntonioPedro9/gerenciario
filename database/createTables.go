package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func CreateDatabaseTables(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL
		);
	`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

