package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vitorsoratto/gojobs/schemas"
)

func sendError(ctx *gin.Context, statusCode int, message string) {
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(statusCode, gin.H{
		"message": message,
	})
}

func sendSuccess(ctx *gin.Context, statusCode int, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(statusCode, gin.H{
		"message": fmt.Sprintf("operation: %s", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
