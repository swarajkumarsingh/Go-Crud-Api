package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/swarajkumarsingh/go-build/config"
	"github.com/swarajkumarsingh/go-build/controllers"
	"github.com/swarajkumarsingh/go-build/db"
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
	r := gin.Default()

	r.GET("/", controllers.HealthCheck)
	r.GET("/users", controllers.GetUsers)
	r.POST("/user", controllers.CreateUser)
	r.GET("/user/:id", controllers.GetUser)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	fmt.Println("Server running on port " + PORT)
	r.Run("127.0.0.1:" + PORT)
}