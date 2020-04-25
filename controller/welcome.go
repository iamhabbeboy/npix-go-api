package controller

import "github.com/gin-gonic/gin"

//Welcome index route
func Welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "pong",
	})
}