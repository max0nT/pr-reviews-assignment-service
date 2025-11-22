package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/max0nT/pr-assign/internal/entities"
	"github.com/rs/zerolog/log"
)

func (cnt *Controllers) GetUsers(ctx *gin.Context) {
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
	res, err := cnt.TeamManage.GetUsers(&userParams)
	if err != nil {
		log.Print("Error during users get: ", err.Error())
		cnt.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, res)

}
