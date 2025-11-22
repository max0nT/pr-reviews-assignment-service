package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/max0nT/pr-assign/internal/entities"
	"github.com/rs/zerolog/log"
)

// OpenPr @Summary Create PR and assign random reviewers
//
//	@Tags		PR manage
//	@Param		request	body		entities.PrCreate	true	"PR create parameters"
//	@Success	201		{object}	entities.PrRead
//	@Failure	400		{object}	entities.RequestError
//	@Failure	404		{object}	entities.RequestError
//	@Failure	500		{objects}	entities.RequestError
//	@Router		/api/v1/pr/open/ [post]
//
// Open PR and assign some reviewers.
func (c *Controllers) OpenPr(ctx *gin.Context) {
	var validatedData entities.PrCreate
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

	res, err := c.PrManage.OpenPr(validatedData)
	if err != nil {
		c.HandleError(ctx, err)
		log.Print("Error during team add: ", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, res)
}
