package main

import (
	"golang-advance/assignment-1/handler"
	"golang-advance/assignment-1/repository/postgres_gorm"
	"golang-advance/assignment-1/router"
	"golang-advance/assignment-1/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	dsn := "postgresql://postgres:postgres@localhost:5432/postgres"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	userRepo := postgres_gorm.NewUserRepository(gormDB)
	submissionRepo := postgres_gorm.NewSubmissionRepository(gormDB)

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	submissionService := service.NewSubmissionService(submissionRepo)
	submissionHandler := handler.NewSubmissionHandler(submissionService)

	router.SetupRouter(r, userHandler, submissionHandler)

	log.Println("Running server on port 8080")
	r.Run(":8080")
}
