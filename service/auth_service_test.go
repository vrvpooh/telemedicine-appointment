package service

import (
	"database/sql"
	"log"
	"telemedicine-api/config"
	"telemedicine-api/model"
	"telemedicine-api/repository"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// เปลี่ยนชื่อเพื่อไม่ให้ซ้ำกับโมดูลอื่น
func setupAuthTestDB() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	config.DB = db

	schema := `
	CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT UNIQUE,
		password TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	config.DB.Exec(schema)
}

func TestRegister(t *testing.T) {
	setupAuthTestDB() // เรียกใช้ฟังก์ชันชื่อใหม่
	repo := &repository.AuthRepository{}
	svc := &AuthService{Repo: repo}

	t.Run("Success Register", func(t *testing.T) {
		req := model.RegisterRequest{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "password123",
		}
		user, err := svc.Register(req)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if user.Email != req.Email {
			t.Errorf("expected email %s, got %s", req.Email, user.Email)
		}
	})

	t.Run("Duplicate Email", func(t *testing.T) {
		req := model.RegisterRequest{
			Name:     "Test User 2",
			Email:    "test@example.com",
			Password: "password123",
		}
		_, err := svc.Register(req)
		if err == nil || err.Error() != "email already exists" {
			t.Errorf("expected 'email already exists' error, got %v", err)
		}
	})
}

func TestLogin(t *testing.T) {
	setupAuthTestDB() // เรียกใช้ฟังก์ชันชื่อใหม่
	repo := &repository.AuthRepository{}
	svc := &AuthService{Repo: repo}

	// Register a user first
	svc.Register(model.RegisterRequest{
		Name:     "Login User",
		Email:    "login@example.com",
		Password: "correct_password",
	})

	t.Run("Success Login", func(t *testing.T) {
		req := model.LoginRequest{
			Email:    "login@example.com",
			Password: "correct_password",
		}
		token, err := svc.Login(req)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if token == "" {
			t.Error("expected token, got empty string")
		}
	})

	t.Run("Wrong Password", func(t *testing.T) {
		req := model.LoginRequest{
			Email:    "login@example.com",
			Password: "wrong_password",
		}
		_, err := svc.Login(req)
		if err == nil || err.Error() != "invalid email or password" {
			t.Errorf("expected invalid credentials error, got %v", err)
		}
	})
}
