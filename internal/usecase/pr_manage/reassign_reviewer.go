package prmanage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

func (pm *PrManage) ReassignUserReviewer( // nolint: cyclop
	reviewerData *entities.PrUnassign,
) (res entities.PrAssign, err error) {
	ctx := context.Background()
	tx, err := pm.Cfg.Pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return
	}
	defer pm.Cfg.CloseTxForFail(ctx, &tx, err)

	pr, err := pm.PrRepo.SelectPr(
		ctx,
		&tx,
		&entities.PrParams{PrId: reviewerData.PrId},
	)
	if err != nil {
		return
	}

	prValidationError := ""
	if len(pr) == 0 {
		prValidationError += fmt.Sprintf(
			"Pr with id %s does not exist",
			reviewerData.PrId,
		)
	}
	if len(pr) != 0 && pr[0].IsMerged {
		prValidationError += fmt.Sprintf(
			"Pr with id %s mustn't be merged to do reassign",
			reviewerData.PrId,
		)
	}
	if prValidationError != "" {
		err = &entities.RequestError{
			Msg:        prValidationError,
			StatusCode: http.StatusBadRequest,
		}
		return
	}

	rmCount, err := pm.PrRepo.DeleteReviewer(ctx, &tx, reviewerData)
	if err != nil {
		return
	}
	if rmCount == 0 {
		err = &entities.RequestError{
			Msg: fmt.Sprintf(
				"There isn't reviewer with id %s for pr %s",
				reviewerData.OldUserId,
				reviewerData.PrId,
			),
			StatusCode: http.StatusBadRequest,
		}
		return
	}

	reviewers, err := pm.PrRepo.SelectReviewer(
		ctx,
		&tx,
		entities.PrReviewerParams{PrId: pr[0].PrId},
	)
	if err != nil {
		return
	}

	prohibitForAssign := []string{}
	prohibitForAssign = append(
		prohibitForAssign,
		pr[0].CreatedBy,
		reviewerData.OldUserId,
	)

	for _, reviewerData := range reviewers {
		prohibitForAssign = append(prohibitForAssign, reviewerData.ReviewerId)
	}
	newReviewer, err := pm.UserRepo.SelectUsers(
		ctx,
		&tx,
		&entities.UserParams{
			NotIdIn:  prohibitForAssign,
			TeamName: pr[0].CreatedByData.TeamName,
			IsActive: true,
			Limit:    1,
		},
	)
	if err != nil {
		return
	}
	if len(newReviewer) == 0 {
		err = &entities.RequestError{
			Msg:        "Not found new reviewer for pr with id " + reviewerData.PrId,
			StatusCode: http.StatusBadRequest,
		}
		return
	}

	err = pm.PrRepo.InsertReviewers(ctx, &tx, &pr[0], &newReviewer)
	if err != nil {
		return
	}

	res.NewUserId = newReviewer[0].Id
	res.PrId = pr[0].PrId

	err = tx.Commit(ctx)
	return
}
