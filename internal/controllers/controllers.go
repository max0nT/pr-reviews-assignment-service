package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/max0nT/pr-assign/internal/usecase"
)

type Controllers struct {
	TeamManage usecase.TeamManage
	Validator  *validator.Validate
}

func New(tm usecase.TeamManage, v *validator.Validate) *Controllers {
	return &Controllers{
		TeamManage: tm,
		Validator:  v,
	}
}
