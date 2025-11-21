package teammanage

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

func (tm *TeamManage) ChangeUserActive(
	userData *entities.UserChangeActive,
) (res entities.User, err error) {
	ctx := context.Background()
	tx, err := tm.Cfg.Pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		tx.Rollback(ctx) // nolint: errcheck, gosec
		return
	}

	res, err = tm.UserRepo.UpdateStatus(ctx, &tx, *userData)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = &entities.RequestError{
				Msg: fmt.Sprintf(
					"User with id %s does not exist",
					userData.Id,
				),
				StatusCode: http.StatusBadRequest,
			}
		}
	}

	tx.Commit(ctx) // nolint: errcheck, gosec
	return
}
