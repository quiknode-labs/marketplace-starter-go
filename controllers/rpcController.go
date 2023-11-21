package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quiknode-labs/marketplace-starter-go/initializers"
	"github.com/quiknode-labs/marketplace-starter-go/models"
	"github.com/quiknode-labs/marketplace-starter-go/types"
)

func RPC(c *gin.Context) {

	// get data off the request body
	var requestBody types.RpcRequest
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"error": "could not parse JSON",
		})
	}

	// get data of the request header
	isTest := c.Request.Header.Get("X-QN-TESTING")
	quicknodeId := c.Request.Header.Get("x-quicknode-id")
	endpointId := c.Request.Header.Get("x-instance-id")
	chain := c.Request.Header.Get("x-qn-chain")
	network := c.Request.Header.Get("x-qn-network")
	log.Println("/rpc with", chain, network, quicknodeId, requestBody.Method)

	// find the account
	var account models.Account
	findAccountResult := initializers.DB.Where("quicknode_id = ?", quicknodeId).First(&account)
	if findAccountResult.Error != nil {
		c.JSON(404, gin.H{
			"error": "could not find account",
		})
		return
	}

	// find the endpoint
	var endpoint models.Endpoint
	findEndpointResult := initializers.DB.Where("account_id = ? AND quicknode_id = ? AND chain = ? AND network = ?", account.ID, endpointId, chain, network).First(&endpoint)
	if findEndpointResult.Error != nil {
		c.JSON(404, gin.H{
			"error": "could not find endpoint",
		})
		return
	}

	requestBodyJsonString, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshalling JSON"})
		return
	}

	// Create and store RpcRequest in database
	rpcRequest := models.RpcRequest{
		AccountId:   account.ID,
		EndpointID:  endpoint.ID,
		MethodName:  requestBody.Method,
		RequestBody: string(requestBodyJsonString),
		Chain:       chain,
		Network:     network,
		IsTest:      isTest == "true",
	}
	rpcRequestSaved := initializers.DB.Create(&rpcRequest)
	if rpcRequestSaved.Error != nil {
		c.JSON(500, gin.H{
			"error": "could not create rpc request",
		})
		return
	}

	// FILLME: ADD YOUR CODE HERE
	// Make sure you set rpcRequestSaved.ResponseStatus and ResponseBody
	rpcRequest.ResponseBody = "{\"abc\": 123}"
	rpcRequest.ResponseStatus = 200
	rpcRequestSaved = initializers.DB.Save(&rpcRequest)
	if rpcRequestSaved.Error != nil {
		c.JSON(500, gin.H{
			"error": "could not update rpc response body and status",
		})
		return
	}

	// FILLME: prepare result to send back
	result := gin.H{
		"quicknode-id": quicknodeId,
		"chain":        chain,
		"network":      network,
		"method":       requestBody.Method,
		"params":       requestBody.Params,
	}

	// Return JSON
	c.JSON(200, gin.H{
		"ID":      requestBody.ID,
		"Jsonrpc": requestBody.JsonRpc,
		"Result":  result,
	})
}
