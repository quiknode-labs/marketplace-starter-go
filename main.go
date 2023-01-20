package main

import(
  "github.com/quiknode-labs/token-dash/initializers"
  "github.com/quiknode-labs/token-dash/controllers"

  "github.com/gin-gonic/gin"
)

func init() {
  initializers.LoadEnvVariables()
  initializers.ConnectToDB()
}

func main() {
  r := gin.Default()

  r.POST("/provision", controllers.Provision)
  r.DELETE("/deprovision", controllers.Deprovision)
  r.PUT("/update", controllers.Update)
  r.DELETE("/deactivate_endpoint", controllers.DeactivateEndpoint)

  r.Run()
}
