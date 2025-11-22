package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/max0nT/pr-assign/internal/entities"
	"github.com/rs/zerolog/log"
)

// MergePr @Summary Mar PR as merged
//
//	@Tags		PR manage
//	@Param		request	body		entities.PrMerge	true	"PR merge parameters"
//	@Success	200		{object}	entities.PrSimple
//	@Failure	400		{object}	entities.RequestError
//	@Failure	404		{object}	entities.RequestError
//	@Failure	500		{objects}	entities.RequestError
//	@Router		/api/v1/pr/merge/ [patch]
//
// Mark Pr as merged.
func (c *Controllers) MergePr(ctx *gin.Context) {
	var validatedData entities.PrMerge
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

	res, err := c.PrManage.MergePr(validatedData)
	if err != nil {
		c.HandleError(ctx, err)
		log.Print("Error during team add: ", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, res)
}
