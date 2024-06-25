package main

import (
	"context"
	"golang-advance/session-7-db-pg-gorm/handler"
	"golang-advance/session-7-db-pg-gorm/repository/postgres_gorm"
	"golang-advance/session-7-db-pg-gorm/repository/postgres_pgx"
	"golang-advance/session-7-db-pg-gorm/router"
	"golang-advance/session-7-db-pg-gorm/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
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

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.SetupRouter(r, userHandler)

	log.Println("Running server on port 8080")
	r.Run(":8080")
}

func connectDB(dbURL string) (postgres_pgx.PgxPoolIface, error) {
	return pgxpool.New(context.Background(), dbURL)
}
