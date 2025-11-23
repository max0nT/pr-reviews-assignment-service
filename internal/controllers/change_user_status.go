package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/max0nT/pr-assign/internal/entities"
	"github.com/rs/zerolog/log"
)

// ChangeUserActive @Summary Change user active status
//
//	@Tags		Team manage
//	@Param		request	body		entities.UserChangeActive	true	"User Change active status parameters"
//	@Success	200		{object}	entities.User
//	@Failure	400		{object}	entities.RequestError
//	@Failure	404		{object}	entities.RequestError
//	@Failure	500		{objects}	entities.RequestError
//	@Router		/api/v1/user/change-status-active/ [patch]
//
// Change user active status.
func (c *Controllers) ChangeUserActive(ctx *gin.Context) {
	var validatedData entities.UserChangeActive

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

	res, err := c.TeamManage.ChangeUserActive(&validatedData)
	if err != nil {
		c.HandleError(ctx, err)
		log.Print("Error during user change status: ", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, res)

}
