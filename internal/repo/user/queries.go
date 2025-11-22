package userrepo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

func (repo *UserRepository) SelectUsers( // nolint: cyclop
	ctx context.Context,
	tx *pgx.Tx,
	userParams *entities.UserParams,
) (res []entities.User, err error) {
	queryBuilder := repo.Cfg.Builder.Select("*").From("users")
	if userParams.Id != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"id": userParams.Id})
	}
	if userParams.NotId != "" {
		queryBuilder = queryBuilder.Where(sq.NotEq{"id": userParams.NotId})
	}
	if len(userParams.NotIdIn) != 0 {
		queryBuilder = queryBuilder.Where(sq.NotEq{"id": userParams.NotIdIn})
	}
	if userParams.TeamName != "" {
		queryBuilder = queryBuilder.Where(
			sq.Eq{"team_name": userParams.TeamName},
		)
	}
	if userParams.IsActive {
		queryBuilder = queryBuilder.Where(
			sq.Eq{"is_active": userParams.IsActive},
		)
	}
	if userParams.Limit != 0 {
		queryBuilder = queryBuilder.Limit(
			uint64(userParams.Limit), // nolint: gosec
		)
	}

	queryString, args, err := queryBuilder.ToSql()
	if err != nil {
		return
	}

	rawRes, err := (*tx).Query(ctx, queryString, args...)
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

func (repo *UserRepository) InsertUsers(
	ctx context.Context,
	tx *pgx.Tx,
	users []entities.User,
	teamName string,
) (res []entities.User, err error) {
	queryBuilder := repo.Cfg.Builder.Insert("users").
		Columns("id", "username", "team_name", "is_active")

	for _, user := range users {
		queryBuilder = queryBuilder.Values(
			user.Id,
			user.Username,
			teamName,
			user.IsActive,
		)
	}

	queryBuilder = queryBuilder.Suffix(`ON CONFLICT (id) DO UPDATE SET
		username = EXCLUDED.username,
		team_name = EXCLUDED.team_name,
		is_active = EXCLUDED.is_active`,
	)
	queryBuilder = queryBuilder.Suffix("RETURNING *")

	queryString, args, err := queryBuilder.ToSql()
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
	if err = rawRes.Err(); err != nil {
		return
	}

	defer rawRes.Close()
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
	return
}

func (repo *UserRepository) UpdateStatus(
	ctx context.Context,
	tx *pgx.Tx,
	userData entities.UserChangeActive,
) (res entities.User, err error) {
	queryBuilder := repo.Cfg.Builder.Update("users").
		Set("is_active", userData.IsActive).
		Where(sq.Eq{"id": userData.Id}).
		Suffix("RETURNING *")

	queryString, args, err := queryBuilder.ToSql()
	if err != nil {
		return
	}

	rawRes := (*tx).QueryRow(
		ctx,
		queryString,
		args...,
	)
	err = rawRes.Scan(
		&res.Id,
		&res.Username,
		&res.TeamName,
		&res.IsActive,
	)

	return
}
