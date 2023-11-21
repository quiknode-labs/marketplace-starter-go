package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	// Setup static files and templates
	r.LoadHTMLGlob("templates/*")
	r.Static("/images", "./images")
	r.Static("/js", "./js")

	// Set up session middleware
	store := cookie.NewStore([]byte("8036e05e78860fbfa87ef3de97d4d899"))
	r.Use(sessions.Sessions("blockbook_session", store))

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

	r.GET("/dash/:id", controllers.Dashboard)
	r.GET("/dash/:id/requests", controllers.RequestsIndex)

	r.GET("/healthz", controllers.Healthcheck)

	r.GET("/healthcheck", controllers.Healthcheck)

	r.Run()
}
