package controllers

import (
	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	// Return JSON
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
