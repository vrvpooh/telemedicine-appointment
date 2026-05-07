package service

import (
	"errors"
	"telemedicine-api/model"
	"telemedicine-api/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("my_super_secret_key") // ในของจริงควรดึงจาก .env

type AuthService struct {
	Repo *repository.AuthRepository
}

func (s *AuthService) Register(req model.RegisterRequest) (*model.User, error) {
	// เช็คว่ามี email นี้หรือยัง
	existingUser, _ := s.Repo.GetUserByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	err = s.Repo.CreateUser(user)
	return user, err
}

func (s *AuthService) Login(req model.LoginRequest) (string, error) {
	user, err := s.Repo.GetUserByEmail(req.Email)
	if err != nil || user == nil {
		return "", errors.New("invalid email or password")
	}

	// เทียบรหัสผ่าน
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// สร้าง JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token หมดอายุใน 3 วัน
	})

	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}

func (s *AuthService) GetProfile(userID int) (*model.User, error) {
	return s.Repo.GetUserByID(userID)
}
