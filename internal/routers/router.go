package routers

import (
	"ecommerce-backend/internal/controller"
	"ecommerce-backend/internal/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("After --> CC")
}

func NewRouter() *gin.Engine {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	// use the middleware
	r.Use(middlewares.AuthMiddleware(), BB(), CC)

	// Define a simple GET endpoint
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", controller.NewPongController().Pong)
		v1.GET("/user", controller.NewUserController().GetUserByID)
		// v1.PUT("/ping", Pong)
		// v1.PATCH("/ping", Pong)
		// v1.DELETE("/ping", Pong)
		// v1.HEAD("/ping", Pong)
		// v1.OPTIONS("/ping", Pong)
	}

	// v2 := r.Group("/api/v2")
	// {
	// 	v2.GET("/ping", Pong)
	// 	v2.PUT("/ping", Pong)
	// 	v2.PATCH("/ping", Pong)
	// 	v2.DELETE("/ping", Pong)
	// 	v2.HEAD("/ping", Pong)
	// 	v2.OPTIONS("/ping", Pong)
	// }
	return r
}
