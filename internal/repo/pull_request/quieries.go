package prrepo

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

func (repo *PrRepository) InsertPr(
	ctx context.Context,
	tx *pgx.Tx,
	insertData *entities.PrCreate,
) (res entities.PrSimple, err error) {
	queryBuilder := repo.Cfg.Builder.Insert("pull_requests").
		Columns("id", "name", "created_by_id").
		Values(insertData.PrId, insertData.PrName, insertData.CreatedBy).
		Suffix("ON CONFLICT (id) DO NOTHING").
		Suffix("RETURNING id, name, created_by_id, is_merged, created_at")

	queryString, args, err := queryBuilder.ToSql()
	if err != nil {
		return
	}

	resRaw := (*tx).QueryRow(ctx, queryString, args...)

	err = resRaw.Scan(
		&res.PrId,
		&res.PrName,
		&res.CreatedBy,
		&res.IsMerged,
		&res.CreatedAt,
	)

	return
}

func (repo *PrRepository) MergePr(
	ctx context.Context,
	tx *pgx.Tx,
	prData *entities.PrMerge,
) (res entities.PrSimple, err error) {
	queryBuilder := repo.Cfg.Builder.Update("pull_requests").
		Set("is_merged", true).
		Set("merged_at", time.Now()).
		Where(sq.Eq{"id": prData.PrId}).
		Suffix("RETURNING *")
	queryString, args, err := queryBuilder.ToSql()
	if err != nil {
		return
	}

	rawRes := (*tx).QueryRow(ctx, queryString, args...)

	err = rawRes.Scan(
		&res.PrId,
		&res.PrName,
		&res.IsMerged,
		&res.CreatedBy,
		&res.CreatedAt,
		&res.MergedAt,
	)

	return
}

func (repo *PrRepository) SelectPr(
	ctx context.Context,
	tx *pgx.Tx,
	prData *entities.PrParams,
) (res []entities.PrSimple, err error) {
	queryBuilder := repo.Cfg.Builder.Select(
		"pull_requests.id",
		"pull_requests.name",
		"pull_requests.is_merged",
		"pull_requests.created_by_id",
		"pull_requests.created_at",
		"pull_requests.merged_at",
		"users.id",
		"users.username",
		"users.team_name",
		"users.is_active",
	).From("pull_requests")
	if prData.PrId != "" {
		queryBuilder = queryBuilder.Where(
			sq.Eq{"pull_requests.id": prData.PrId},
		)
	}
	if prData.CreatedBy != "" {
		queryBuilder = queryBuilder.Where(
			sq.Eq{"created_by_id": prData.CreatedBy},
		)
	}
	if prData.IsMerged {
		queryBuilder = queryBuilder.Where(sq.Eq{"is_merged": prData.IsMerged})
	}
	if prData.ReviewerId != "" {
		queryBuilder = queryBuilder.LeftJoin(
			"reviewers ON pull_requests.id = reviewers.pr_id",
		).Where(sq.Eq{"reviewers.reviewer_id": prData.ReviewerId})
	}
	queryBuilder = queryBuilder.Join(
		"users ON pull_requests.created_by_id = users.id",
	)

	queryString, args, err := queryBuilder.ToSql()
	if err != nil {
		return
	}

	rawRes, err := (*tx).Query(ctx, queryString, args...)
	if err != nil {
		return
	}

	for rawRes.Next() {
		var prData entities.PrSimple

		err = rawRes.Scan(
			&prData.PrId,
			&prData.PrName,
			&prData.IsMerged,
			&prData.CreatedBy,
			&prData.CreatedAt,
			&prData.MergedAt,
			&prData.CreatedByData.Id,
			&prData.CreatedByData.Username,
			&prData.CreatedByData.TeamName,
			&prData.CreatedByData.IsActive,
		)
		if err != nil {
			return
		}

		res = append(res, prData)
	}

	rawRes.Close()
	return
}

func (repo *PrRepository) SelectReviewer(
	ctx context.Context,
	tx *pgx.Tx,
	reviewerData entities.PrReviewerParams,
) (res []entities.PrReviewer, err error) {
	queryBuilder := repo.Cfg.Builder.Select("reviewer_id", "pr_id").
		From("reviewers")

	queryString, args, err := queryBuilder.ToSql()
	if err != nil {
		return
	}

	rawRes, err := (*tx).Query(ctx, queryString, args...)
	if err != nil {
		return
	}

	for rawRes.Next() {
		var reviewerData entities.PrReviewer

		err = rawRes.Scan(
			&reviewerData.ReviewerId,
			&reviewerData.PrId,
		)
		if err != nil {
			return
		}

		res = append(res, reviewerData)
	}

	return
}

func (repo *PrRepository) DeleteReviewer(
	ctx context.Context,
	tx *pgx.Tx,
	prData *entities.PrUnassign,
) (res int64, err error) {
	queryBuilder := repo.Cfg.Builder.Delete("reviewers").
		Where(sq.Eq{"reviewer_id": prData.OldUserId, "pr_id": prData.PrId})

	queryString, args, err := queryBuilder.ToSql()
	if err != nil {
		return
	}

	rawRes, err := (*tx).Exec(ctx, queryString, args...)
	if err != nil {
		return
	}
	res = rawRes.RowsAffected()
	return
}

func (repo *PrRepository) InsertReviewers(
	ctx context.Context,
	tx *pgx.Tx,
	prData *entities.PrSimple,
	userData *[]entities.User,
) (err error) {
	queryBuilder := repo.Cfg.Builder.Insert("reviewers").
		Columns("reviewer_id", "pr_id")
	for _, user := range *userData {
		queryBuilder = queryBuilder.Values(user.Id, prData.PrId)
	}

	queryString, args, err := queryBuilder.ToSql()
	if err != nil {
		return
	}

	_, err = (*tx).Exec(ctx, queryString, args...)

	return
}
