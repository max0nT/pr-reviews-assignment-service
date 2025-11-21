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
