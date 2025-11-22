package prmanage

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

func (pm *PrManage) GetPr(
	prParams *entities.PrParams,
) (res []entities.PrSimple, err error) {
	ctx := context.Background()
	tx, err := pm.Cfg.Pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return
	}
	defer pm.Cfg.CloseTxForFail(ctx, &tx, err)

	res, err = pm.PrRepo.SelectPr(ctx, &tx, prParams)

	if err == nil {
		err = tx.Commit(ctx)
	}

	return

}
