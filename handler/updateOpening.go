package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorsoratto/gojobs/schemas"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	req := UpdateOpeningRequest{}
	ctx.BindJSON(&req)

	if err := req.Validate(); err != nil {
		logger.Errorf("Error validating opening request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "queryParameter [id] is required")
		return
	}

	dbOpening := schemas.Opening{}

	if err := db.First(&dbOpening, id).Error; err != nil {
		e := fmt.Sprintf("opening with id [%v] not found", id)
		logger.Errorf(e)
		sendError(ctx, http.StatusNotFound, e)
		return
	}

	if req.Role != "" {
		dbOpening.Role = req.Role
	}

	if req.Company != "" {
		dbOpening.Company = req.Company
	}

	if req.Location != "" {
		dbOpening.Location = req.Location
	}

	if req.Link != "" {
		dbOpening.Link = req.Link
	}

	if req.Description != "" {
		dbOpening.Description = req.Description
	}

	if req.Salary > 0 {
		dbOpening.Salary = req.Salary
	}

	if req.Remote != nil {
		dbOpening.Remote = *req.Remote
	}

	if err := db.Save(&dbOpening).Error; err != nil {
		e := fmt.Sprintf("Error while updating opening with id [%v]", id)
		logger.Errorf("%v | error: %v", e, err)
		sendError(ctx, http.StatusInternalServerError, e)
		return
	}

	sendSuccess(ctx, http.StatusOK, "update-opening", dbOpening)
}
