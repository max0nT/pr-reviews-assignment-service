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
		InsertUsers(
			context.Context,
			*pgx.Tx,
			[]entities.User,
			string,
		) ([]entities.User, error)
	}
)
