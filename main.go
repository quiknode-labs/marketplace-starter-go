package main

import (
	"net/http"
	"os"

	"github.com/quiknode-labs/marketplace-starter-go/controllers"
	"github.com/quiknode-labs/marketplace-starter-go/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "<h1>marketplace-starter-go</h1>")
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		os.Getenv("BASIC_AUTH_USERNAME"): os.Getenv("BASIC_AUTH_PASSWORD"),
	}))

	authorized.POST("/provision", controllers.Provision)
	authorized.DELETE("/deprovision", controllers.Deprovision)
	authorized.PUT("/update", controllers.Update)
	authorized.DELETE("/deactivate_endpoint", controllers.DeactivateEndpoint)

	r.POST("/rpc", controllers.RPC)

	r.GET("/api", controllers.API)

	r.GET("/dashboard", controllers.Dashboard)

	r.GET("/healthcheck", controllers.Healthcheck)

	r.Run()
}
