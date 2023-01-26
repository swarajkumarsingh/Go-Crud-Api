package errorhandler

import (
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleErrorWithMsg(msg ...string) {
	if len(msg) == 0 {
		msg[0] = "Something went wrong"
	}
	panic(msg[0])
}

func Recovery(c *gin.Context, httpStatusCode int) {
	if r := recover(); r != nil {
		log.Println("recovered: ", r)
		msg, ok := r.(string)

		if ok {
			// string message passed - no need to report to sentry
			CustomError(c, msg, httpStatusCode)
		} else {
			err, ok := r.(error)
			msg = "something went wrong, please check back in an hour"
			if ok {
				// report to sentry (other services)
				CustomErrorSentry(c, httpStatusCode, msg, err)
			} else {
				// when string or error cannot be recovered (rare case)
				CustomError(c, "something went wrong, please check back in an hour", 500)
			}
		}
	}
}

func CustomError(c *gin.Context, msg string, httpStatusCode int) {
	c.JSON(httpStatusCode, gin.H{
		"error": msg,
	})
}

func CustomErrorSentry(c *gin.Context, httpStatusCode int, msg string, err error) {
	if os.Getenv("Stage") != "PROD" {
		// report to sentry first if environment is prod, uat or dev
		ReportToSentry(c, err)
		CustomError(c, msg, httpStatusCode)
	}
	CustomError(c, msg, httpStatusCode)
}

func ReportToSentry(context *gin.Context, err error) {
	hub := sentry.GetHubFromContext(context)
	hub.CaptureException(err)
}
