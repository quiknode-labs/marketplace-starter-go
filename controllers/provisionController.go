package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/quiknode-labs/qn-go-add-on/initializers"
	"github.com/quiknode-labs/qn-go-add-on/models"
)

func Provision(c *gin.Context) {
	// Get data off request body
	var requestBody struct {
		QuicknodeID string `json:"quicknode-id"`
		Plan        string `json:"plan"`
		EndpointId  string `json:"endpoint-id"`
		WssUrl      string `json:"wss-url"`
		HttpUrl     string `json:"http-url"`
		Chain       string `json:"chain"`
		Network     string `json:"network"`
	}

	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"error": "could not parse JSON",
		})
	}
	log.Println("/provision with", requestBody)

	// Create an account
	account := models.Account{QuicknodeID: requestBody.QuicknodeID, Plan: requestBody.Plan}
	accountResult := initializers.DB.Create(&account)
	if accountResult.Error != nil {
		c.JSON(500, gin.H{
			"error": "could not create account",
		})
		return
	}

	// Create an endpoint
	endpoint := models.Endpoint{
		QuicknodeID: requestBody.EndpointId,
		WssUrl:      requestBody.WssUrl,
		HttpUrl:     requestBody.HttpUrl,
		Chain:       requestBody.Chain,
		Network:     requestBody.Network,
	}
	endpoint.AccountID = account.ID
	endpointResult := initializers.DB.Create(&endpoint)
	if endpointResult.Error != nil {
		c.JSON(500, gin.H{
			"error": "could not create endpoint",
		})
		return
	}

	scheme := "http://"
	if c.Request.TLS != nil {
		scheme = "https://"
	}

	// Return JSON
	c.JSON(200, gin.H{
		"status":        "success",
		"dashboard-url": scheme + c.Request.Host + "/dashboard",
		"access-url":    "",
	})
}

func Update(c *gin.Context) {
	// Get data off request body
	var requestBody struct {
		QuicknodeID string `json:"quicknode-id"`
		Plan        string `json:"plan"`
		EndpointId  string `json:"endpoint-id"`
		WssUrl      string `json:"wss-url"`
		HttpUrl     string `json:"http-url"`
		Chain       string `json:"chain"`
		Network     string `json:"network"`
	}

	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"error": "could not parse JSON",
		})
	}
	log.Println("/update with", requestBody)

	// find the account
	var account models.Account
	findAccountResult := initializers.DB.Where("quicknode_id = ?", requestBody.QuicknodeID).First(&account)
	if findAccountResult.Error != nil {
		c.JSON(404, gin.H{
			"error": "could not find account",
		})
		return
	}
	account.Plan = requestBody.Plan
	updateAccountResult := initializers.DB.Save(&account)
	if updateAccountResult.Error != nil {
		c.JSON(500, gin.H{
			"error": "could not update account",
		})
		return
	}

	// find the endpoint
	var endpoint models.Endpoint
	findEndpointResult := initializers.DB.Where("quicknode_id = ?", requestBody.EndpointId).First(&endpoint)
	if findEndpointResult.Error != nil {
		c.JSON(404, gin.H{
			"error": "could not find endpoint",
		})
		return
	}
	endpoint.WssUrl = requestBody.WssUrl
	endpoint.HttpUrl = requestBody.HttpUrl
	endpoint.Chain = requestBody.Chain
	endpoint.Network = requestBody.Network
	updateEndpointResult := initializers.DB.Save(&endpoint)
	if updateEndpointResult.Error != nil {
		c.JSON(500, gin.H{
			"error": "could not update endpoint",
		})
		return
	}

	// Return JSON
	c.JSON(200, gin.H{
		"status": "success",
	})
}

func DeactivateEndpoint(c *gin.Context) {
	// Get data off request body
	var requestBody struct {
		QuicknodeID string `json:"quicknode-id"`
		EndpointId  string `json:"endpoint-id"`
		Chain       string `json:"chain"`
		Network     string `json:"network"`
	}

	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"error": "could not parse JSON",
		})
	}
	log.Println("/deactivate_endpoint with", requestBody)

	// find the account
	var account models.Account
	findAccountResult := initializers.DB.Where("quicknode_id = ?", requestBody.QuicknodeID).First(&account)
	if findAccountResult.Error != nil {
		c.JSON(404, gin.H{
			"error": "could not find account",
		})
		return
	}

	// find the endpoint
	var endpoint models.Endpoint
	findEndpointResult := initializers.DB.Where("quicknode_id = ?", requestBody.EndpointId).First(&endpoint)
	if findEndpointResult.Error != nil {
		c.JSON(404, gin.H{
			"error": "could not find endpoint",
		})
		return
	}

	// delete the endpoint
	deleteResult := initializers.DB.Delete(&endpoint)
	if deleteResult.Error != nil {
		c.JSON(500, gin.H{
			"error": "could not deactive endpoint",
		})
		return
	}

	// Return JSON
	c.JSON(200, gin.H{
		"status": "success",
	})
}

func Deprovision(c *gin.Context) {
	// Get data off request body
	var requestBody struct {
		QuicknodeID string `json:"quicknode-id"`
	}

	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"error": "could not parse JSON",
		})
	}
	log.Println("/deprovision with", requestBody)

	// find the account
	var account models.Account
	findResult := initializers.DB.Where("quicknode_id = ?", requestBody.QuicknodeID).First(&account)
	if findResult.Error != nil {
		c.JSON(404, gin.H{
			"error": "could not find account",
		})
		return
	}

	// delete the account
	deleteResult := initializers.DB.Delete(&account)
	if deleteResult.Error != nil {
		c.JSON(500, gin.H{
			"error": "could not deprovision account",
		})
		return
	}

	// Return JSON
	c.JSON(200, gin.H{
		"status": "success",
	})
}
