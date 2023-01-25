package main

import(
  "github.com/quiknode-labs/token-dash/initializers"
  "github.com/quiknode-labs/token-dash/models"
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
