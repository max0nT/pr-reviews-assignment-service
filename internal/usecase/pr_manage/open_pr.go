package prmanage

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

const (
	DefaultReviewersCount int = 2
)

func (pm *PrManage) OpenPr( // nolint: cyclop
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
	defer pm.Cfg.CloseTx(ctx, &tx, &err)

	createdBy, err := pm.UserRepo.SelectUsers(
		ctx,
		&tx,
		&entities.UserParams{Id: prData.CreatedBy, IsActive: true},
	)
	if err != nil {
		return
	}

	CreatedByValidateErrorMessage := ""
	if len(createdBy) == 0 {
		CreatedByValidateErrorMessage += fmt.Sprintf(
			"User with id %s does not exist \n",
			prData.CreatedBy,
		)
	}
	if len(createdBy) != 0 && !createdBy[0].IsActive {
		CreatedByValidateErrorMessage += fmt.Sprintf(
			"User with id %s does not exist \n",
			prData.CreatedBy,
		)
	}
	if CreatedByValidateErrorMessage != "" {
		err = &entities.RequestError{
			Msg:        CreatedByValidateErrorMessage,
			StatusCode: http.StatusBadRequest,
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
		return
	}
	if len(reviewers) == 0 {
		return
	}

	err = pm.PrRepo.InsertReviewers(
		ctx,
		&tx,
		&insertedPr,
		&reviewers,
	)
	if err != nil {
		return
	}
	res.PrId = insertedPr.PrId
	res.PrName = insertedPr.PrName
	res.CreatedBy = insertedPr.CreatedBy
	res.CreatedAt = insertedPr.CreatedAt
	res.IsMerged = insertedPr.IsMerged
	res.MergedAt = *insertedPr.MergedAt
	res.Reviewers = reviewers

	return
}
