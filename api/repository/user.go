package repository

import (
	"context"
	"golang-boilerplate/ent"
	"golang-boilerplate/internal/db"
	dto "golang-boilerplate/internal/types"
)

type UserRepository interface {
	CreateUser(input *dto.UserCreateInput) (*ent.User, error)
	GetAllUsers() ([]*ent.User, error)
}

type UserRepositoryImpl struct {
	Client *ent.UserClient
}

func NewUserRepository() UserRepository {
	userClient := db.GetClient().User

	return &UserRepositoryImpl{
		Client: userClient,
	}
}

func (u *UserRepositoryImpl) CreateUser(input *dto.UserCreateInput) (*ent.User, error) {
	user, err := u.Client.Create().
		SetID(input.ID).
		SetName(input.Name).
		Save(context.TODO())
	return user, err
}

func (u *UserRepositoryImpl) GetAllUsers() ([]*ent.User, error) {
	user, err := u.Client.Query().All(context.TODO())
	return user, err
}
