package service

import (
	"context"
	"golang-advance/assignment-1/entity"
)

type IUserService interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

type userService struct {
	userRepo IUserRepository
}

func NewUserService(userRepo IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	panic("implement me")
}

func (u *userService) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	panic("implement me")
}

func (u *userService) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	panic("implement me")
}

func (u *userService) DeleteUser(ctx context.Context, id int) error {
	panic("implement me")
}

func (u *userService) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	panic("implement me")
}
