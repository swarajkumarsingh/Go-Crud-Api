package errorhandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery(c *gin.Context) {
	if r := recover(); r != nil {
		log.Println("recovered: ", r)
		msg, ok := r.(string)

		if ok {
			// string message passed - no need to report to sentry
			c.JSON(http.StatusOK, gin.H{
				"error": msg,
			})
		} else {
			_, ok := r.(error)
			if ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "Something went wrong",
				})
			}
		}
	}
}
