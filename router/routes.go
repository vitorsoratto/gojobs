package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vitorsoratto/gojobs/handler"
)

func initRoutes(router *gin.Engine) {
	api := router.Group("/api/")
	{
		api.GET("/openings", handler.ListOpeningsHandler)
		api.GET("/opening", handler.GetOpeningHandler)
		api.POST("/opening", handler.CrateOpeningHandler)
		api.PUT("/opening", handler.UpdateOpeningHandler)
		api.DELETE("/opening", handler.DeleteOpeningHandler)
	}
}
