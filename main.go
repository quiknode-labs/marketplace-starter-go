package main

import (
	"os"

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

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		os.Getenv("BASIC_AUTH_USERNAME"): os.Getenv("BASIC_AUTH_PASSWORD"),
	}))

	authorized.POST("/provision", controllers.Provision)
	authorized.DELETE("/deprovision", controllers.Deprovision)
	authorized.PUT("/update", controllers.Update)
	authorized.DELETE("/deactivate_endpoint", controllers.DeactivateEndpoint)

	r.POST("/rpc", controllers.RPC)

	r.GET("/dashboard", controllers.Dashboard)

	r.GET("/healthcheck", controllers.Healthcheck)

	r.Run()
}
