package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/max0nT/pr-assign/internal/entities"
	"github.com/rs/zerolog/log"
)

// ReassignReviewer @Summary Change one reviewer to random another one
//
//	@Tags		PR manage
//	@Param		request	body		entities.PrUnassign	true	"Reviewer unassign parameters"
//	@Success	200		{object}	entities.PrAssign
//	@Failure	400		{object}	entities.RequestError
//	@Failure	404		{object}	entities.RequestError
//	@Failure	500		{objects}	entities.RequestError
//	@Router		/api/v1/pr/reassign/ [patch]
//
// Unassign user and replace it by another one.
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
