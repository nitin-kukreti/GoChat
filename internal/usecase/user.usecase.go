package usecase

import "github.com/nitin-kukreti/GoChat/internal/domain"

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) CreateUser(name string) (domain.User, error) {
	return uc.repo.CreateUser(name)
}

func (uc *UserUseCase) GetUserByID(id int) (domain.User, error) {
	return uc.repo.GetUserByID(id)
}
