package teammanage

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

func (tm *TeamManage) AddTeam(
	teamData entities.ItemCreate,
) (res entities.ItemRead, err error) {
	ctx := context.Background()
	tx, err := tm.Cfg.Pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return
	}
	defer tm.Cfg.CloseTxForFail(ctx, &tx, err)

	insertedTeam, err := tm.TeamRepo.InsertTeam(ctx, &tx, teamData.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = &entities.RequestError{
				Msg: fmt.Sprintf(
					"Team with %s already exists",
					teamData.Name,
				),
				StatusCode: http.StatusBadRequest,
			}
		}
		return
	}

	insertedUsers, err := tm.UserRepo.InsertUsers(
		ctx,
		&tx,
		teamData.Users,
		insertedTeam.Name,
	)
	if err != nil {
		return
	}
	res.Id = insertedTeam.Id
	res.Name = insertedTeam.Name
	res.Users = insertedUsers

	err = tx.Commit(ctx)
	return
}
