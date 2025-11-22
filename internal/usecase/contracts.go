package usecase

import "github.com/max0nT/pr-assign/internal/entities"

type (
	TeamManage interface {
		AddTeam(entities.TeamCreate) (entities.TeamRead, error)
		ChangeUserActive(
			*entities.UserChangeActive,
		) (entities.User, error)
		GetUsers(*entities.UserParams) ([]entities.User, error)
		GetUsersStats(*entities.UserParams) ([]entities.UserStats, error)
		GetTeam(*entities.TeamGetParams) (entities.TeamRead, error)
	}
	PrManage interface {
		OpenPr(entities.PrCreate) (entities.PrRead, error)
		MergePr(entities.PrMerge) (entities.PrSimple, error)
		ReassignUserReviewer(
			*entities.PrUnassign,
		) (entities.PrAssign, error)
		GetPr(*entities.PrParams) ([]entities.PrSimple, error)
	}
)
