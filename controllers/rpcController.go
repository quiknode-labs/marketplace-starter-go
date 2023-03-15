package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/quiknode-labs/qn-go-add-on/initializers"
	"github.com/quiknode-labs/qn-go-add-on/models"
	rpc "github.com/quiknode-labs/qn-go-add-on/types"
)

func RPC(c *gin.Context) {

	// get data off the request body
	var requestBody rpc.RpcRequest
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"error": "could not parse JSON",
		})
	}

	// get data of the request header
	quicknodeId := c.Request.Header.Get("x-quicknode-id")
	chain := c.Request.Header.Get("x-qn-chain")
	network := c.Request.Header.Get("x-qn-network")
	log.Println("/rpc with", chain, network, quicknodeId, requestBody.Method)

	// TODO: check if quicknodeId, endpoint-id, chain and network are valid

	// Create and store RpcRequest in database
	rpcRequest := models.RpcRequest{
		QuicknodeID: quicknodeId,
		Chain:       chain,
		Network:     network,
		RequestID:   requestBody.ID.(string),
		Method:      requestBody.Method,
		Version:     requestBody.JsonRpc,
	}
	rpcRequestSaved := initializers.DB.Create(&rpcRequest)
	if rpcRequestSaved.Error != nil {
		c.JSON(500, gin.H{
			"error": "could not create rpc request",
		})
		return
	}

	// FILLME: ADD YOUR CODE HERE

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
