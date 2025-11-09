package handler

import (
	"net/http"

	"github.com/H0wZy/go.learnings.restapi/schemas"
	"github.com/gin-gonic/gin"
)

// ListOpeningsHandler
// @Summary ListOpeningsHandler
// @Description List all jobs openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	var openings []schemas.Opening
	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("error listing openings: %v", err.Error())
		responseError(ctx, http.StatusInternalServerError, false, err.Error())
		return
	}
	responseSuccess(ctx, "list-openings", openings, true)
}
