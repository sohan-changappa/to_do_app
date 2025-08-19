package main

import (
	"gin-auth-sql/database"
	"gin-auth-sql/handlers"
	"gin-auth-sql/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	protected := r.Group("/protect")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/todos", handlers.CreateTodo)
		protected.GET("/todos", handlers.GetTodos)
		protected.PUT("/todos/:id", handlers.UpdateTodo)
		protected.DELETE("/todos/:id", handlers.DeleteTodo)
	}

	r.Run(":8080")
}
