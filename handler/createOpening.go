package handler

import (
	"net/http"

	"github.com/H0wZy/go.learnings.restapi/schemas"
	"github.com/gin-gonic/gin"
)

// CreateOpeningHandler
// @Summary CreateOpeningHandler
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {

	request := CreateOpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		responseError(ctx, http.StatusBadRequest, false, "malformed request JSON")
		return
	} else if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		responseError(ctx, http.StatusBadRequest, false, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		responseError(ctx, http.StatusInternalServerError, false, "error creating opening on database")
		return
	}

	responseSuccess(ctx, "create-opening", opening, true)
	logger.Infof("request received: %+v", request)
}
