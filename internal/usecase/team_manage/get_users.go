package teammanage

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

func (tm *TeamManage) GetUsers(
	userParams *entities.UserParams,
) (res []entities.User, err error) {
	ctx := context.Background()
	tx, err := tm.Cfg.Pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return
	}
	defer tm.Cfg.CloseTxForFail(ctx, &tx, err)

	res, err = tm.UserRepo.SelectUsers(ctx, &tx, userParams)

	if err == nil {
		err = tx.Commit(ctx)
	}

	return
}
