package usecases

import "golang-solid-clean-architecture/entities"

type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id int64) (*entities.User, error)
}

type UserUsecase struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) Register(user *entities.User) error {
	return u.repo.Create(user)
}

func (u *UserUsecase) GetUser(id int64) (*entities.User, error) {
	return u.repo.GetByID(id)
}
