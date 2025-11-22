package teamrepo

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

func (repo *TeamRepository) SelectTeam(
	ctx context.Context,
	tx *pgx.Tx,
	teamData *entities.TeamGetParams,
) (res entities.TeamSimple, err error) {
	queryBuilder := repo.Cfg.Builder.Select("*").
		From("teams").
		Where(sq.Eq{"name": teamData.TeamName})

	queryString, args, err := queryBuilder.ToSql()
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

func (repo *TeamRepository) InsertTeam(
	ctx context.Context,
	tx *pgx.Tx,
	teamName string,
) (res entities.TeamSimple, err error) {
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
