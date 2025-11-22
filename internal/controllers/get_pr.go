package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/max0nT/pr-assign/internal/entities"
	"github.com/rs/zerolog/log"
)

// GetPr @Summary Get team by name
//
//	@Tags		PR manage
//	@Param		request	query		entities.PrParams	false	"PR search parameters"
//	@Success	200		{array}		entities.PrSimple
//	@Failure	500		{objects}	entities.RequestError
//	@Router		/api/v1/pr/ [get]
//
// Get PR list by search params.
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
