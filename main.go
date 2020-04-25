package main

import (
	// "os"
	"log"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/iamhabbeboy/api/controller"
	"github.com/iamhabbeboy/api/database"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	database.CategorySeeder()
	database.PictureSeeder()
	database.GameSeeder()

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		// v1.POST("/", createTodo)
		v1.GET("/", controller.Welcome)
		v1.GET("/category", controller.FetchCategory)
		v1.GET("/game/:id", controller.FetchGame)
		// v1.GET("/:id", fetchSingleTodo)
		// v1.PUT("/:id", updateTodo)
		// v1.DELETE("/:id", deleteTodo)
	}
	router.Run()
}
