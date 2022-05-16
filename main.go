package main

import (
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.RecordUaAndTime)

	blogEngine := engine.Group("/article")
	{
		v1 := blogEngine.Group("/v1")
		{
			v1.POST("/add", controller.BlogAdd)
			// v1.GET("/list", controller.BlogList)
			// v1.PUT("/list", controller.BlogUpdate)
			// v1.DELETE("/delete", controller.BlogDelete)
		}
	}
	engine.Run(":23000")
}