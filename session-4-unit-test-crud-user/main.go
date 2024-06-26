package main

import (
	"golang-advance/session-4-unit-test-crud-user/entity"
	"golang-advance/session-4-unit-test-crud-user/handler"
	"golang-advance/session-4-unit-test-crud-user/repository/slice"
	"golang-advance/session-4-unit-test-crud-user/router"
	"golang-advance/session-4-unit-test-crud-user/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	var mockUserDBInSlice []entity.User
	userRepo := slice.NewUserRepository(mockUserDBInSlice)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.SetupRouter(r, userHandler)

	log.Println("Running server on port 8080")
	r.Run(":8080")
}
