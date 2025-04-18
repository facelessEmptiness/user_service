package usecase

import (
	"github.com/facelessEmptiness/user_service/internal/domain"
	"github.com/facelessEmptiness/user_service/internal/repository"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (u *UserUseCase) Register(user *domain.User) (string, error) {
	return u.repo.Create(user)
}

func (u *UserUseCase) GetByEmail(email string) (*domain.User, error) {
	return u.repo.GetByEmail(email)
}

func (u *UserUseCase) GetByID(id string) (*domain.User, error) {
	return u.repo.GetByID(id)
}
