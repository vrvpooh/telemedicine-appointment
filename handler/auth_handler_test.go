package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"telemedicine-api/config"
	"telemedicine-api/model"
	"telemedicine-api/repository"
	"telemedicine-api/service"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func setupHandlerTest() *gin.Engine {
	db, _ := sql.Open("sqlite3", ":memory:")
	config.DB = db
	config.DB.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, email TEXT UNIQUE, password TEXT, created_at DATETIME);")

	repo := &repository.AuthRepository{}
	svc := &service.AuthService{Repo: repo}
	SetupAuthHandler(svc)

	r := gin.Default()
	r.POST("/api/auth/register", RegisterUser)
	r.POST("/api/auth/login", LoginUser)
	return r
}

func TestRegisterHandler(t *testing.T) {
	r := setupHandlerTest()

	t.Run("Valid Registration", func(t *testing.T) {
		body := model.RegisterRequest{
			Name:     "Handler Test",
			Email:    "handler@test.com",
			Password: "password123",
		}
		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(jsonBody))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("expected 201, got %d", w.Code)
		}
	})
}
