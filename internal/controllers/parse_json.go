package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func (c *Controllers) ParseJson(
	ctx *gin.Context,
	modelData any,
) (err error) {
	rawData, err := ctx.GetRawData()

	if err != nil {
		return
	}

	err = json.Unmarshal(rawData, &modelData)
	return
}
