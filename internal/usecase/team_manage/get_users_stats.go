package teammanage

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

func (tm *TeamManage) GetUsersStats(
	userParams *entities.UserParams,
) (res []entities.UserStats, err error) {
	ctx := context.Background()
	tx, err := tm.Cfg.Pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return
	}
	defer tm.Cfg.CloseTx(ctx, &tx, &err)

	res, err = tm.UserRepo.SelectUsersStats(ctx, &tx, userParams)

	return
}
