package controllers

import (
	"github.com/gin-gonic/gin"
)

func API(c *gin.Context) {
	// FILLME: implement API authentication here with an API key or something
	c.JSON(200, gin.H{
		"status": "ok",
		"data":   "your api data goes here",
	})
}
