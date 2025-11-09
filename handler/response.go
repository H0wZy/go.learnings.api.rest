package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func responseError(ctx *gin.Context, code int, success bool, msg string) {
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.JSON(code, gin.H{
		"statusCode": code,
		"success":    success,
		"message":    msg,
	})
}
func responseSuccess(ctx *gin.Context, opr string, data interface{}, success bool) {
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"success":    success,
		"message":    fmt.Sprintf("operation from handler: '%s' completed successfully", opr),
		"data":       data,
	})
}
