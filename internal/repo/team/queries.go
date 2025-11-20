package teamrepo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

func (repo *ItemRepository) InsertTeam(
	ctx context.Context,
	tx *pgx.Tx,
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

	resRaw := (*tx).QueryRow(ctx, queryString, args...)

	err = resRaw.Scan(
		&res.Id,
		&res.Name,
	)

	return
}
