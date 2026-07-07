package auth

import (
	"database/sql"
	"errors"

	"queueflow/internal/models"
	"queueflow/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	users *repository.UserRepository
}

func NewAuthService(
	db *sql.DB,
) *AuthService {

	return &AuthService{
		users: repository.NewUserRepository(db),
	}
}

func (s *AuthService) Register(
	username string,
	email string,
	password string,
) error {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := models.User{
		ID:       uuid.New().String(),
		Username: username,
		Email:    email,
		Password: string(hash),
		Role:     "user",
	}

	return s.users.Create(user)
}

func (s *AuthService) Login(
	email string,
	password string,
) (string, error) {

	user, err := s.users.FindByEmail(email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	s.users.UpdateLastLogin(user.ID)

	token, err := GenerateToken(
		user.ID,
		user.Role,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
