package main

import (
	"golang-advance/session-5-validator/entity"
	"golang-advance/session-5-validator/handler"
	"golang-advance/session-5-validator/repository/slice"
	"golang-advance/session-5-validator/router"
	"golang-advance/session-5-validator/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// setup service
	var mockUserDBInSlice []entity.User
	userRepo := slice.NewUserRepository(mockUserDBInSlice)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routes
	router.SetupRouter(r, userHandler)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
