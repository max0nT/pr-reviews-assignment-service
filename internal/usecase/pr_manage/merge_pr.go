package prmanage

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

func (pm *PrManage) MergePr(
	prData entities.PrMerge,
) (res entities.PrSimple, err error) {
	ctx := context.Background()
	tx, err := pm.Cfg.Pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return
	}
	defer pm.Cfg.CloseTxForFail(ctx, &tx, err)

	res, err = pm.PrRepo.MergePr(
		ctx,
		&tx,
		&prData,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = &entities.RequestError{
				Msg: fmt.Sprintf(
					"Pr with id %s does not exist",
					prData.PrId,
				),
				StatusCode: http.StatusNotFound,
			}
		}
		return
	}
	err = tx.Commit(ctx)

	return
}
