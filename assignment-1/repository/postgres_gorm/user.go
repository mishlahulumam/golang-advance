package postgres_gorm

import (
	"context"
	"golang-advance/assignment-1/entity"
	"golang-advance/assignment-1/service"
)

type userRepository struct {
	db GormDBIface
}

func NewUserRepository(db GormDBIface) service.IUserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	panic("implement me")
}

func (u *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	panic("implement me")
}

func (u *userRepository) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	panic("implement me")
}

func (u *userRepository) DeleteUser(ctx context.Context, id int) error {
	panic("implement me")
}

func (u *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	panic("implement me")
}
