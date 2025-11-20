package teamrepo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

func (repo *ItemRepository) InsertTeam(
	teamName string,
) (res entities.ItemSimple, err error) {
	queryBuild := repo.Cfg.Builder.Insert("teams").
		Columns("name").
		Values(teamName).
		Suffix("ON CONFLICT (name) DO NOTHING").
		Suffix("RETURNING *")

	queryString, args, err := queryBuild.ToSql()
	if err != nil {
		return
	}

	tx, err := repo.Cfg.Pool.BeginTx(
		context.Background(),
		pgx.TxOptions{
			IsoLevel: pgx.RepeatableRead,
		},
	)
	if err != nil {
		return
	}

	resRaw := tx.QueryRow(context.Background(), queryString, args...)

	err = resRaw.Scan(
		&res.Id,
		&res.Name,
	)
	if err != nil {
		tx.Rollback(context.Background()) // nolint: errcheck, gosec
		return
	}
	tx.Commit(context.Background()) //nolint: errcheck, gosec

	return
}
