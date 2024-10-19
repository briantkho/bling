package user

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUserByID(userID uuid.UUID) (*User, error) {
	return s.repo.GetUserByID(userID)
}

func (s *Service) CreateUser(firstName string, lastName string, username string, email string, password string) (uuid.UUID, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPasswordString := string(hashedPassword)

	if err != nil {
		return uuid.Nil, err
	}

	user := &User{
		FirstName:           firstName,
		LastName:            lastName,
		Email:               email,
		Username:            username,
		Password:            hashedPasswordString,
		CreatedAt:           time.Now(),
		UserID:              uuid.New(),
		CurrentLevelPoints:  0,
		TotalLifetimePoints: 0,
		LevelID:             1,
	}
	userId, err := s.repo.CreateUser(user)
	if err != nil {
		return uuid.Nil, err
	}

	return userId, nil
}
