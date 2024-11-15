package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorsoratto/gojobs/schemas"
)

// @BasePath /api
// @Summary Create opening
// @Accept json
// @Produce json
// @Param opening body CreateOpeningRequest true "opening"
// @Success 201 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	req := CreateOpeningRequest{}

	ctx.BindJSON(&req)

	if err := req.Validate(); err != nil {
		logger.Errorf("Error validating opening request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:        req.Role,
		Company:     req.Company,
		Location:    req.Location,
		Remote:      *req.Remote,
		Link:        req.Link,
		Description: req.Description,
		Salary:      req.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error while creating opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Error while creating opening on database")
		return
	}

	sendSuccess(ctx, http.StatusCreated, "create-opening", opening)
}
