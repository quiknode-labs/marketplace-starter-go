package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/quiknode-labs/marketplace-starter-go/initializers"
	"github.com/quiknode-labs/marketplace-starter-go/models"
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
	isTest := c.Request.Header.Get("X-QN-TESTING")

	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"error": "could not parse JSON",
		})
	}
	log.Println("/provision with", requestBody)

	// find the account
	var account models.Account
	findAccountResult := initializers.DB.Where("quicknode_id = ?", requestBody.QuicknodeID).First(&account)
	if findAccountResult.Error != nil {
		// Create an account since we did not find one
		isTest := c.Request.Header.Get("X-QN-TESTING")
		account = models.Account{
			QuicknodeID: requestBody.QuicknodeID,
			PlanSlug:    requestBody.Plan,
			IsTest:      isTest == "true",
		}
		accountResult := initializers.DB.Create(&account)
		if accountResult.Error != nil {
			c.JSON(500, gin.H{
				"error": "could not create account",
			})
			return
		}
		log.Println(" --> created account", account.ID)
	} else {
		log.Println(" --> found account", account.ID)
	}

	// Create an endpoint
	endpoint := models.Endpoint{
		AccountID:   account.ID,
		QuicknodeID: requestBody.EndpointId,
		WssUrl:      requestBody.WssUrl,
		HttpUrl:     requestBody.HttpUrl,
		Chain:       requestBody.Chain,
		Network:     requestBody.Network,
		IsTest:      isTest == "true",
	}
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
		"dashboard-url": scheme + c.Request.Host + "/dash/" + strconv.Itoa(int(account.ID)),
		"access-url":    scheme + c.Request.Host + "/api", // Note: should be protected by API key
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
	account.PlanSlug = requestBody.Plan
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
	isTest := c.Request.Header.Get("X-QN-TESTING")
	endpoint.IsTest = isTest == "true"
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
