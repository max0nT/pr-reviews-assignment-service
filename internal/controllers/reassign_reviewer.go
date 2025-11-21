package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/max0nT/pr-assign/internal/entities"
	"github.com/rs/zerolog/log"
)

func (c *Controllers) ReassignReviewer(ctx *gin.Context) {
	var validatedData entities.PrUnassign
	if err := c.ParseJson(ctx, &validatedData); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			entities.RequestError{
				Msg:        err.Error(),
				StatusCode: http.StatusBadRequest,
			},
		)
		return
	}

	if err := c.Validator.Struct(validatedData); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			entities.RequestError{
				Msg:        err.Error(),
				StatusCode: http.StatusBadRequest,
			},
		)
		return
	}

	res, err := c.PrManage.ReassignUserReviewer(&validatedData)
	if err != nil {
		c.HandleError(ctx, err)
		log.Print("Error during reassign reviewer: ", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, res)
}
