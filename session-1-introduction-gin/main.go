package main

import (
	"golang-advance/session-1-introduction-gin/middleware"
	"golang-advance/session-1-introduction-gin/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Instalasi router Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(middleware.AuthMiddleware())

	// Definisikan route
	router.SetupRouter(r)

	// Jalankan server pada port 8080
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
