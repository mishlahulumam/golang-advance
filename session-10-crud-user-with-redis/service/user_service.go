package service

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-advance/session-10-crud-user-with-redis/entity"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
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
	rdb      *redis.Client
}

const redisUserByIDKey = "user:%d"

func NewUserService(userRepo IUserRepository, rdb *redis.Client) IUserService {
	return &userService{userRepo: userRepo, rdb: rdb}
}

func (s *userService) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	createdUser, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("gagal membuat pengguna: %v", err)
	}

	fmt.Println(createdUser)
	redisKey := fmt.Sprintf(redisUserByIDKey, createdUser.ID)
	createdUserJSON, err := json.Marshal(createdUser)
	if err != nil {
		log.Println("gagal marshal json")
	}
	if err := s.rdb.Set(ctx, redisKey, createdUserJSON, 60*time.Second).Err(); err != nil {
		log.Println("error when set redis key", redisKey)
	}
	return createdUser, nil
}

func (s *userService) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	redisKey := fmt.Sprintf(redisUserByIDKey, id)
	val, err := s.rdb.Get(ctx, redisKey).Result()
	if err == nil {
		log.Println("data tersedia di redis")
		err = json.Unmarshal([]byte(val), &user)
		if err != nil {
			log.Println("gagal marshall data di redis, coba ambil ke database")
		}
	}
	if err != nil {
		log.Println("data tidak tersedia di redis, ambil dari database")
		user, err = s.userRepo.GetUserByID(ctx, id)
		if err != nil {
			log.Println("gagal ambil data di database")
			return entity.User{}, fmt.Errorf("gagal mendapatkan pengguna berdasarkan ID: %v", err)
		}
	}

	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	updatedUser, err := s.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("gagal memperbarui pengguna: %v", err)
	}
	return updatedUser, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int) error {
	redisKey := fmt.Sprintf(redisUserByIDKey, id)
	if err := s.rdb.Del(ctx).Err(); err != nil {
		log.Println("gagal delete key redis", redisKey)
	}
	err := s.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus pengguna: %v", err)
	}
	return nil
}

func (s *userService) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	users, err := s.userRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan semua pengguna: %v", err)
	}
	return users, nil
}
