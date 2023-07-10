package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func loadEnvVariables() error {
	err := godotenv.Load("../.env")
	if err != nil {
		return err
	}

	return nil
}

func CreateDatabaseConnection() (*sql.DB, error) {
	err := loadEnvVariables()
	if err != nil {
		return nil, err
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user, password, host, port, dbname)

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
	err := loadEnvVariables()
	if err != nil {
		return nil, err
	}

	user := os.Getenv("TEST_DB_USER")
	password := os.Getenv("TEST_DB_PASSWORD")
	host := os.Getenv("TEST_DB_HOST")
	port := os.Getenv("TEST_DB_PORT")
	dbname := os.Getenv("TEST_DB_NAME")

	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user, password, host, port, dbname)

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
