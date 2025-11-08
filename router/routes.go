package router

import (
	"github.com/H0wZy/go.learnings.restapi/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.CreateOpeningHandler)
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.GET("/opening", handler.DeleteOpeningHandler)
		v1.GET("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}
}
