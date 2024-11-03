package services

import (
	"golang-boilerplate/api/repository"
	"golang-boilerplate/ent"
)

type UserService interface {
	GetAllUsers() ([]*ent.User, error)
	CreateUser()
	UpdateUser()
	DeleteUser()
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService() UserService {
	return &UserServiceImpl{repository.NewUserRepository()}
}

func (u *UserServiceImpl) GetAllUsers() ([]*ent.User, error) {
	users, err := u.userRepository.GetAllUsers()
	return users, err
}
func (u *UserServiceImpl) CreateUser() {}
func (u *UserServiceImpl) UpdateUser() {}
func (u *UserServiceImpl) DeleteUser() {}
