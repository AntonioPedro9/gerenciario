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

func (ur *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (name, email, password, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := ur.db.QueryRow(query, user.Name, user.Email, user.Password, time.Now()).Scan(&user.ID)

	if err != nil {
		log.Println("Erro ao criar o usuário:", err)
		return err
	}

	return nil
}

func (ur *UserRepository) List() ([]*models.User, error) {
	query := `
		SELECT id, name, email, password, created_at
		FROM users
	`

	rows, err := ur.db.Query(query)

	if err != nil {
		log.Println("Erro ao buscar os usuários:", err)
		return nil, err
	}

	defer rows.Close()

	users := []*models.User{}

	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

		if err != nil {
			log.Println("Erro ao ler os usuários:", err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) Update(user *models.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, password = $3
		WHERE id = $4
	`

	_, err := ur.db.Exec(query, user.Name, user.Email, user.Password, user.ID)

	if err != nil {
		log.Println("Erro ao atualizar o usuário:", err)
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(id int) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := ur.db.Exec(query, id)

	if err != nil {
		log.Println("Erro ao excluir o usuário:", err)
		return err
	}

	return nil
}
