package controller

import (
	"log"
	"net/http"
	"github.com/iamhabbeboy/api/database"
	"github.com/iamhabbeboy/api/model"
	"github.com/gin-gonic/gin"
)
//FetchGame retrieve all game with category
func FetchGame(c *gin.Context) {
	categoryID := c.Param("id")
	if categoryID == "" {
		log.Fatal("Category ID not found")
		c.JSON(http.StatusNotFound, gin.H{"result": "Invalid category ID",})
		return
	}
	db := database.Connect()
	var game []model.Game
	query := db.Find(&game, categoryID)
	defer query.Close()

	if len(game) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"result": "Games not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": query,})
}