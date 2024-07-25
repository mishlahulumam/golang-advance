package postgres_gorm_raw

import (
	"context"
	"errors"
	"golang-advance/session-10-crud-user-with-redis/entity"
	"golang-advance/session-10-crud-user-with-redis/service"
	"log"

	"gorm.io/gorm"
)

type GormDBIface interface {
	WithContext(ctx context.Context) *gorm.DB
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}

type userRepository struct {
	db GormDBIface
}

func NewUserRepository(db GormDBIface) service.IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	query := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id"
	var createdID int
	if err := r.db.WithContext(ctx).
		Raw(query, user.Name, user.Email, user.Password).Scan(&createdID).Error; err != nil {
		log.Printf("Error creating user: %v\n", err)
		return entity.User{}, err
	}
	user.ID = createdID
	return *user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1"
	if err := r.db.WithContext(ctx).Raw(query, id).Scan(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, nil
		}
		log.Printf("Error getting user by ID: %v\n", err)
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	query := "UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3"
	if err := r.db.WithContext(ctx).Exec(query, user.Name, user.Email, id).Error; err != nil {
		log.Printf("Error updating user: %v\n", err)
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	if err := r.db.WithContext(ctx).Exec(query, id).Error; err != nil {
		log.Printf("Error deleting user: %v\n", err)
		return err
	}
	return nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	query := "SELECT id, name, email, password, created_at, updated_at FROM users"
	var users []entity.User
	if err := r.db.WithContext(ctx).Raw(query).Scan(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, nil
		}
		log.Printf("Error getting all users: %v\n", err)
		return nil, err
	}
	return users, nil
}
