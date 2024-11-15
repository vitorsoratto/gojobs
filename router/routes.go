package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/vitorsoratto/gojobs/docs"
	"github.com/vitorsoratto/gojobs/handler"
)

func initRoutes(router *gin.Engine) {
	handler.InitHandler()
	basePath := "/api"
	docs.SwaggerInfo.BasePath = basePath

	api := router.Group(basePath)
	{
		api.GET("/openings", handler.ListOpeningsHandler)
		api.GET("/opening", handler.GetOpeningHandler)
		api.POST("/opening", handler.CreateOpeningHandler)
		api.PUT("/opening", handler.UpdateOpeningHandler)
		api.DELETE("/opening", handler.DeleteOpeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
