package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swarajkumarsingh/go-build/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/", controllers.HealthCheck)
	incomingRoutes.GET("/users", controllers.GetUsers)
	incomingRoutes.POST("/user", controllers.CreateUser)
	incomingRoutes.GET("/user/:id", controllers.GetUser)
	incomingRoutes.PUT("/user/:id", controllers.UpdateUser)
	incomingRoutes.DELETE("/user/:id", controllers.DeleteUser)
}
