package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST",
	})
}


