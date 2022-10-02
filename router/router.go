package router

import (
	"Assignment3/handler"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/weather")
	{
		userRouter.GET("/status", handler.GetStatusHandler)
	}

	return r
}
