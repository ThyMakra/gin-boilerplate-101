package utils

import (
	"net/http"

	"github.com/ThyMakra/gin-boilerplate/backend/schemas"
	"github.com/gin-gonic/gin"
)

func ApiResponse(ctx *gin.Context, message string, code int, data interface{}) {
	jsonResponse := schemas.ReponseSchema{
		StatusCode: code,
		Message:    message,
		Data:       data,
	}

	if code >= 400 {
		ctx.AbortWithStatusJSON(code, jsonResponse)
	} else {
		ctx.JSON(code, jsonResponse)
	}
}

func ErrorResponse(ctx *gin.Context, error interface{}) {
	err := schemas.SchemaErrorResponse{
		StatusCode: http.StatusBadRequest,
		Error:      error,
	}

	ctx.AbortWithStatusJSON(err.StatusCode, err)
}
