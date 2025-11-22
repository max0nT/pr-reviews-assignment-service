package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/max0nT/pr-assign/internal/entities"
	"github.com/rs/zerolog/log"
)

// AddTeam @Summary Add team with members
//
//	@Tags		Team manage
//	@Param		request	body		entities.TeamCreate	true	"Team add parameters"
//	@Success	201		{object}	entities.TeamRead
//	@Failure	400		{object}	entities.RequestError
//	@Failure	404		{object}	entities.RequestError
//	@Failure	500		{objects}	entities.RequestError
//	@Router		/api/v1/team/add/ [post]
//
// Add team with members.
func (c *Controllers) AddTeam(ctx *gin.Context) {
	var validatedData entities.TeamCreate
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

	res, err := c.TeamManage.AddTeam(validatedData)
	if err != nil {
		c.HandleError(ctx, err)
		log.Print("Error during team add: ", err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, res)
}
