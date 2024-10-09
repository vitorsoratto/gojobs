package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorsoratto/gojobs/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("Error while listing openings: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Error while listing openings on database")
		return
	}

	sendSuccess(ctx, http.StatusOK, "list-openings", openings)
}
