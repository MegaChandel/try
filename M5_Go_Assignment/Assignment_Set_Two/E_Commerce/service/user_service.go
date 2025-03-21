package service

import (
	"ecommerce-inventory/model"
	"ecommerce-inventory/repository"
	"errors"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) RegisterUser(user *model.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("invalid user data")
	}

	return service.repo.RegisterUser(user)
}

func (service *UserService) AuthenticateUser(username, password string) (*model.User, error) {
	user, err := service.repo.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if user.Password != password {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}
