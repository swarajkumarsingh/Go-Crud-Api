package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errorhandler "github.com/swarajkumarsingh/go-build/errorHandler"
)

func HealthCheck(c *gin.Context) {
	defer errorhandler.Recovery(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "Health OK",
	})
}
