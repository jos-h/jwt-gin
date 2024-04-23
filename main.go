package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jwt-gin/controllers"
	"github.com/jwt-gin/middlewares"
	"github.com/jwt-gin/models"
)

func init() {
	models.SetupDB()
}

func main() {
	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/users", controllers.GetUsers)
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8000")
}
