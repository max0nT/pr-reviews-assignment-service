package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/max0nT/pr-assign/internal/entities"
	"github.com/rs/zerolog/log"
)

// GetTeam @Summary Get team by name
//
//	@Tags		Team manage
//	@Param		request	query		entities.TeamGetParams	false	"Team search parameters"
//	@Success	200		{object}	entities.TeamRead
//	@Failure	404		{object}	entities.RequestError
//	@Failure	500		{objects}	entities.RequestError
//	@Router		/api/v1/team/ [get]
//
// Get Team Using search params (optional).
func (c *Controllers) GetTeam(ctx *gin.Context) {
	var teamParams entities.TeamGetParams
	if err := ctx.BindQuery(&teamParams); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			entities.RequestError{
				Msg:        err.Error(),
				StatusCode: http.StatusBadRequest,
			},
		)
		return
	}

	if err := c.Validator.Struct(teamParams); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			entities.RequestError{
				Msg:        err.Error(),
				StatusCode: http.StatusBadRequest,
			},
		)
		return
	}

	res, err := c.TeamManage.GetTeam(&teamParams)
	if err != nil {
		log.Print("Error during team get: " + err.Error())
		c.HandleError(ctx, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}
