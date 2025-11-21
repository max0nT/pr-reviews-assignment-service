package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/max0nT/pr-assign/internal/usecase"
)

type Controllers struct {
	TeamManage usecase.TeamManage
	PrManage   usecase.PrManage
	Validator  *validator.Validate
}

func New(
	tm usecase.TeamManage,
	pm usecase.PrManage,
	v *validator.Validate,
) *Controllers {
	return &Controllers{
		TeamManage: tm,
		PrManage:   pm,
		Validator:  v,
	}
}
