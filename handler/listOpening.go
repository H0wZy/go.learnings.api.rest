package handler

import (
	"net/http"

	"github.com/H0wZy/go.learnings.restapi/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	var openings []schemas.Opening
	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("error listing openings: %v", err.Error())
		responseError(ctx, http.StatusInternalServerError, false, err.Error())
		return
	}
	responseSuccess(ctx, "list-openings", openings, true)
}
