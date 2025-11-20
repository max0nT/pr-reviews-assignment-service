package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/max0nT/pr-assign/internal/entities"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

const internalMessageError string = "Internal error"
const internalErrorStatusCode int = 500

func (cnt *Controllers) HandleError(
	ctx *gin.Context,
	err error,
) {
	message := internalMessageError
	statusCode := internalErrorStatusCode

	reqErr, ok := err.(*entities.RequestError) // nolint: errorlint
	if ok {
		message = reqErr.Error()
		statusCode = reqErr.StatusCode
	}

	ctx.JSON(
		statusCode,
		ErrorMessage{
			message,
		},
	)
}
