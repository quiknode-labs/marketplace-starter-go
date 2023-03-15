package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/quiknode-labs/qn-go-add-on/initializers"
)

func Healthcheck(c *gin.Context) {
	// make a call to DB to make sure it's up
	var count int64
	initializers.DB.Table("accounts").Count(&count)

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
