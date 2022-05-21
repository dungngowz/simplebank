package util

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ErrorBadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
}

func ErrorInternalServer(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
}

func ErrorNotFound(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusNotFound, ErrorResponse(err))
}

func ErrorUnauthorized(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusUnauthorized, ErrorResponse(err))
}

func ErrorResponse(err error) gin.H {
	var validationErrors validator.ValidationErrors

	writeToLog(err.Error())

	if errors.As(err, &validationErrors) {
		formatedErrs := make([]ErrorMsg, len(validationErrors))
		for i, fe := range validationErrors {
			formatedErrs[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
		}

		return gin.H{
			"errors": formatedErrs,
		}
	}

	return gin.H{
		"error": err.Error(),
	}
}

func writeToLog(msg string) {
	logFile, logErr := os.OpenFile("system.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if logErr != nil {
		log.Fatal(logErr)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.Println(msg)
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "email":
		return "This field must be a valid email address"
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "min":
		return "This field must be at least " + fe.Param() + " characters"
	}
	return "Unknown error"
}
