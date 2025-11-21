package prmanage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

const (
	DefaultReviewersCount int = 2
)

func (pm *PrManage) OpenPr(
	prData entities.PrCreate,
) (res entities.PrRead, err error) {
	ctx := context.Background()

	tx, err := pm.Cfg.Pool.BeginTx(
		ctx,
		pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		},
	)
	if err != nil {
		return
	}

	createdBy, err := pm.UserRepo.SelectUsers(
		ctx,
		&tx,
		&entities.UserParams{Id: prData.CreatedBy, IsActive: true},
	)
	if err != nil {
		tx.Rollback(ctx) // nolint: errcheck, gosec
		return
	}
	if len(createdBy) == 0 {
		tx.Rollback(ctx) // nolint: errcheck, gosec
		err = &entities.RequestError{
			Msg: fmt.Sprintf(
				"User with id %s does not exist",
				prData.CreatedBy,
			),
		}
		return
	}

	insertedPr, err := pm.PrRepo.InsertPr(
		ctx,
		&tx,
		&prData,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = &entities.RequestError{
				Msg: fmt.Sprintf(
					"Pr with id %s already exist",
					prData.PrId,
				),
			}

		}
		tx.Rollback(ctx) // nolint: errcheck, gosec
		return
	}

	reviewers, err := pm.UserRepo.SelectUsers(
		ctx,
		&tx,
		&entities.UserParams{
			NotId:    createdBy[0].Id,
			TeamName: createdBy[0].TeamName,
			IsActive: true,
			Limit:    DefaultReviewersCount,
		},
	)
	if err != nil {
		tx.Rollback(ctx) // nolint: errcheck, gosec
		return
	}
	if len(reviewers) == 0 {
		tx.Commit(ctx) // nolint: errcheck, gosec
		return
	}

	err = pm.PrRepo.InsertReviewers(
		ctx,
		&tx,
		&insertedPr,
		&reviewers,
	)
	if err != nil {
		tx.Rollback(ctx) // nolint: errcheck, gosec
		return
	}

	res.PrId = insertedPr.PrId
	res.PrName = insertedPr.PrName
	res.CreatedBy = insertedPr.CreatedBy
	res.CreatedAt = insertedPr.CreatedAt
	res.IsMerged = insertedPr.IsMerged
	res.MergedAt = insertedPr.MergedAt
	res.Reviewers = reviewers

	err = tx.Commit(ctx) // nolint: errcheck, gosec
	return
}
