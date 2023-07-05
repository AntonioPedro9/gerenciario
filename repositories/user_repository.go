package repositories

import (
	"database/sql"
	"log"
	"time"

	"server/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Create(user *models.User) (*models.User, error) {
	query := `
		INSERT INTO users (name, email, password, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := ur.db.QueryRow(query, user.Name, user.Email, user.Password, time.Now()).Scan(&user.ID)
	if err != nil {
		log.Println("Failed to create a new user:", err)
		return nil, err
	}

	log.Println("New user created:", user.ID)
	return user, nil
}

func (ur *UserRepository) List() ([]*models.User, error) {
	query := `
		SELECT id, name, email, password, created_at
		FROM users
	`

	rows, err := ur.db.Query(query)
	if err != nil {
		log.Println("Failed to fetch users:", err)
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}

	for rows.Next() {
		user := &models.User{}

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			log.Println("Failed to scan user row:", err)
			return nil, err
		}

		users = append(users, user)
	}

	log.Println("Fetched", len(users), "users")
	return users, nil
}

func (ur *UserRepository) GetUserById(id int) (*models.User, error) {
	query := `
		SELECT id, name, email, password, created_at
		FROM users
		WHERE id = $1
		LIMIT 1
	`

	row := ur.db.QueryRow(query, id)
	user := &models.User{}

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println("Failed to fetch user by email:", err)
		return nil, err
	}

	log.Println("Fetched user by ID:", id)
	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, name, email, password, created_at
		FROM users
		WHERE email = $1
		LIMIT 1
	`

	row := ur.db.QueryRow(query, email)
	user := &models.User{}

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println("Failed to fetch user by email:", err)
		return nil, err
	}

	log.Println("Fetched user by email:", email)
	return user, nil
}

func (ur *UserRepository) Update(user *models.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, password = $3
		WHERE id = $4
	`

	_, err := ur.db.Exec(query, user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		log.Println("Failed to update user:", err)
		return err
	}

	log.Println("User updated:", user.ID)
	return nil
}

func (ur *UserRepository) Delete(id int) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := ur.db.Exec(query, id)
	if err != nil {
		log.Println("Failed to delete user:", err)
		return err
	}

	log.Println("User deleted:", id)
	return nil
}
