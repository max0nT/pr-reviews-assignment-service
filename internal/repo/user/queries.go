package userrepo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

func (repo *UserRepository) InsertUsers(
	ctx context.Context,
	tx *pgx.Tx,
	users []entities.User,
	teamName string,
) (res []entities.User, err error) {
	queryBuild := repo.Cfg.Builder.Insert("users").
		Columns("id", "username", "team_name", "is_active")

	for _, user := range users {
		queryBuild = queryBuild.Values(
			user.Id,
			user.Username,
			teamName,
			user.IsActive,
		)
	}

	queryBuild = queryBuild.Suffix(`ON CONFLICT (id) DO UPDATE SET
		username = EXCLUDED.username,
		team_name = EXCLUDED.team_name,
		is_active = EXCLUDED.is_active`,
	)
	queryBuild = queryBuild.Suffix("RETURNING *")

	queryString, args, err := queryBuild.ToSql()
	if err != nil {
		return
	}

	rawRes, err := (*tx).Query(
		ctx,
		queryString,
		args...,
	)
	if err != nil {
		return
	}

	for rawRes.Next() {
		var userData entities.User

		err = rawRes.Scan(
			&userData.Id,
			&userData.Username,
			&userData.TeamName,
			&userData.IsActive,
		)
		if err != nil {
			return
		}
		res = append(res, userData)
	}

	rawRes.Close()
	return
}
