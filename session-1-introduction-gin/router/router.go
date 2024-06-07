package router

import (
	"golang-advance/session-1-introduction-gin/handler"
	"golang-advance/session-1-introduction-gin/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/", handler.RootHandler)

	privateEndpoint := r.Group("/private")
	privateEndpoint.Use(middleware.AuthMiddleware())
	{
		privateEndpoint.POST("/post", handler.PostHandler)
	}
	r.POST("post", handler.PostHandler)
}
