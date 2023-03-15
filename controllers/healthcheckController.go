package controllers

import (
	"github.com/gin-gonic/gin"
)

func Healthcheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
