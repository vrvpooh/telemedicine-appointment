package repository

import (
	"database/sql"
	"telemedicine-api/config"
	"telemedicine-api/model"
)

type AuthRepository struct{}

func (r *AuthRepository) CreateUser(user *model.User) error {
	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`
	result, err := config.DB.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err == nil {
		user.ID = int(id)
	}
	return err
}

func (r *AuthRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	query := `SELECT id, name, email, password, created_at FROM users WHERE email = ?`
	err := config.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil // ไม่พบ user
	}
	return &user, err
}

func (r *AuthRepository) GetUserByID(id int) (*model.User, error) {
	var user model.User
	query := `SELECT id, name, email, created_at FROM users WHERE id = ?`
	// ไม่ต้องดึง password มาตอน get profile ก็ได้
	err := config.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}
