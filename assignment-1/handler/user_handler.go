package handler

import (
	"golang-advance/assignment-1/service"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
}

type UserHandler struct {
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	panic("implement me")
}

func (u *UserHandler) GetUser(c *gin.Context) {
	panic("implement me")
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	panic("implement me")
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	panic("implement me")
}

func (u *UserHandler) GetAllUsers(c *gin.Context) {
	panic("implement me")
}
