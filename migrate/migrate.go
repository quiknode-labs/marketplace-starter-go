package main

import (
	"github.com/quiknode-labs/marketplace-starter-go/initializers"
	"github.com/quiknode-labs/marketplace-starter-go/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Account{})
	initializers.DB.AutoMigrate(&models.Endpoint{})
	initializers.DB.AutoMigrate(&models.RpcRequest{})
}
