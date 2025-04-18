package repository

import (
	"github.com/facelessEmptiness/user_service/internal/domain"
)

type UserRepository interface {
	Create(user *domain.User) (string, error)
	GetByEmail(email string) (*domain.User, error)
	GetByID(id string) (*domain.User, error)
}
