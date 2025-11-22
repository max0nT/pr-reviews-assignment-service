package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

type (
	TeamRepository interface {
		SelectTeam(
			context.Context,
			*pgx.Tx,
			*entities.TeamGetParams,
		) (entities.TeamSimple, error)
		InsertTeam(
			context.Context,
			*pgx.Tx,
			string,
		) (entities.TeamSimple, error)
	}

	UserRepository interface {
		SelectUsers(
			context.Context,
			*pgx.Tx,
			*entities.UserParams,
		) ([]entities.User, error)
		InsertUsers(
			context.Context,
			*pgx.Tx,
			[]entities.User,
			string,
		) ([]entities.User, error)
		UpdateStatus(
			context.Context,
			*pgx.Tx,
			entities.UserChangeActive,
		) (entities.User, error)
	}

	PrRepository interface {
		SelectPr(
			context.Context,
			*pgx.Tx,
			*entities.PrParams,
		) ([]entities.PrSimple, error)
		SelectReviewer(
			context.Context,
			*pgx.Tx,
			entities.PrReviewerParams,
		) ([]entities.PrReviewer, error)
		InsertPr(
			context.Context,
			*pgx.Tx,
			*entities.PrCreate,
		) (entities.PrSimple, error)
		MergePr(
			context.Context,
			*pgx.Tx,
			*entities.PrMerge,
		) (entities.PrSimple, error)
		DeleteReviewer(
			context.Context,
			*pgx.Tx,
			*entities.PrUnassign,
		) (int64, error)
		InsertReviewers(
			context.Context,
			*pgx.Tx,
			*entities.PrSimple,
			*[]entities.User,
		) error
	}
)
