package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/max0nT/pr-assign/internal/entities"
	"github.com/rs/zerolog/log"
)

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
