package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

type (
	TeamRepository interface {
		InsertTeam(
			context.Context,
			*pgx.Tx,
			string,
		) (entities.ItemSimple, error)
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
	}

	PrRepository interface {
		InsertPr(
			context.Context,
			*pgx.Tx,
			*entities.PrCreate,
		) (entities.PrSimple, error)
		InsertReviewers(
			context.Context,
			*pgx.Tx,
			*entities.PrSimple,
			*[]entities.User,
		) error
	}
)
