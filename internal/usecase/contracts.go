package usecase

import "github.com/max0nT/pr-assign/internal/entities"

type (
	TeamManage interface {
		AddTeam(entities.ItemCreate) (entities.ItemRead, error)
		ChangeUserActive(
			*entities.UserChangeActive,
		) (entities.User, error)
		GetUsers(*entities.UserParams) ([]entities.User, error)
	}
	PrManage interface {
		OpenPr(entities.PrCreate) (entities.PrRead, error)
		MergePr(entities.PrMerge) (entities.PrSimple, error)
		ReassignUserReviewer(
			*entities.PrUnassign,
		) (entities.PrAssign, error)
	}
)
