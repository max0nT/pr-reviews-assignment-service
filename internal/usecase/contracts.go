package usecase

import "github.com/max0nT/pr-assign/internal/entities"

type (
	TeamManage interface {
		AddTeam(entities.ItemCreate) (entities.ItemRead, error)
	}
	PrManage interface {
		OpenPr(entities.PrCreate) (entities.PrRead, error)
		MergePr(entities.PrMerge) (entities.PrSimple, error)
	}
)
