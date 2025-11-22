package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/max0nT/pr-assign/internal/entities"
	"github.com/rs/zerolog/log"
)

func (c *Controllers) GetPr(ctx *gin.Context) {
	var prParams entities.PrParams
	err := ctx.BindQuery(&prParams)
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
	res, err := c.PrManage.GetPr(&prParams)
	if err != nil {
		log.Print("Error during users get: ", err.Error())
		c.HandleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, res)

}
