package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorsoratto/gojobs/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "queryParameter [id] is required")
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening).Error; err != nil {
		e := fmt.Sprintf("opening with id [%v] not found", id)
		logger.Errorf(e)
		sendError(ctx, http.StatusNotFound, e)
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		e := fmt.Sprintf("Error while deleting opening with id [%v]", id)
		logger.Errorf("%v | error: %v", e, err)
		sendError(ctx, http.StatusInternalServerError, e)
		return
	}

	sendSuccess(ctx, http.StatusOK, "delete-opening", opening)

}
