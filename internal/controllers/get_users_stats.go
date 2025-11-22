package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/max0nT/pr-assign/internal/entities"
	"github.com/rs/zerolog/log"
)

// GetUsersStats @Summary Get user stats list
//
//	@Tags		Team manage
//	@Param		request	query		entities.UserParams	false	"User search parameters"
//	@Success	200		{array}		entities.UserStats
//	@Failure	400		{object}	entities.RequestError
//	@Failure	500		{objects}	entities.RequestError
//	@Router		/api/v1/user/stats/ [get].
//
// Get User PR statistic.
func (cnt *Controllers) GetUsersStats(ctx *gin.Context) {
	var userParams entities.UserParams
	err := ctx.BindQuery(&userParams)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			entities.RequestError{
				Msg:        err.Error(),
				StatusCode: http.StatusBadRequest,
			},
		)
		return
	}
	res, err := cnt.TeamManage.GetUsersStats(&userParams)
	if err != nil {
		log.Print("Error during users get: ", err.Error())
		cnt.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, res)

}
