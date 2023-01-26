package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/swarajkumarsingh/go-build/config"
	"github.com/swarajkumarsingh/go-build/db"
	"github.com/swarajkumarsingh/go-build/routes"
)

func init() {
	config.LoadEnvVariables()
}

func main() {

	PORT := os.Getenv("PORT")

	// Remove Warning on console
	gin.SetMode(gin.ReleaseMode)

	// Connect to BD
	db.ConnectDB()

	// App Routes
	r := gin.New()
	r.Use(gin.Logger())

	routes.UserRoutes(r)

	fmt.Println("Server running on port " + PORT)
	r.Run("127.0.0.1:" + PORT)
}
