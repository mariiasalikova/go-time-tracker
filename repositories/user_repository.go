// repositories/user_repository.go
package repositories

import (
	"database/sql"
	"errors"
	"github.com/mariiasalikova/go-time-tracker/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (passport_number, surname, name, patronymic, address) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := repo.DB.QueryRow(query, user.PassportNumber, user.Surname, user.Name, user.Patronymic, user.Address).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetUser(id int) (*models.User, error) {
	var user models.User
	query := `SELECT id, passport_number, surname, name, patronymic, address FROM users WHERE id = $1`
	err := repo.DB.QueryRow(query, id).Scan(&user.ID, &user.PassportNumber, &user.Surname, &user.Name, &user.Patronymic, &user.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) UpdateUser(id int, user *models.User) error {
	query := `UPDATE users SET passport_number = $1, surname = $2, name = $3, patronymic = $4, address = $5 WHERE id = $6`
	_, err := repo.DB.Exec(query, user.PassportNumber, user.Surname, user.Name, user.Patronymic, user.Address, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := repo.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) ListUsers() ([]models.User, error) {
	query := `SELECT id, passport_number, surname, name, patronymic, address FROM users`
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.PassportNumber, &user.Surname, &user.Name, &user.Patronymic, &user.Address); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
