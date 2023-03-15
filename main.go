package main

import (
	"github.com/quiknode-labs/qn-go-add-on/controllers"
	"github.com/quiknode-labs/qn-go-add-on/initializers"

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

	r.POST("/rpc", controllers.RPC)

	r.GET("/healthcheck", controllers.Healthcheck)

	r.Run()
}
