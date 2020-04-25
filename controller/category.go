package controller

import (
	"github.com/iamhabbeboy/api/database"
	"github.com/iamhabbeboy/api/model"
	"github.com/gin-gonic/gin"
)

//FetchCategory index route
func FetchCategory(c *gin.Context) {
	db := database.Connect()
	var categories []model.Category
	query := db.Find(&categories)

	defer query.Close()

	c.JSON(200, gin.H{
		"result": query,
	})
}