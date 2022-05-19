package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func ErrorBadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
}
