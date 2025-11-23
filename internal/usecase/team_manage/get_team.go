package teammanage

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

func (tm *TeamManage) GetTeam(
	teamData *entities.TeamGetParams,
) (res entities.TeamRead, err error) {
	ctx := context.Background()

	tx, err := tm.Cfg.Pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return
	}
	defer tm.Cfg.CloseTx(ctx, &tx, &err)

	selectedTeam, err := tm.TeamRepo.SelectTeam(ctx, &tx, teamData)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = &entities.RequestError{
				Msg: fmt.Sprintf(
					"Team with name %s does not exist",
					teamData.TeamName,
				),
				StatusCode: http.StatusNotFound,
			}
		}
		return
	}

	selectedUsers, err := tm.UserRepo.SelectUsers(
		ctx,
		&tx,
		&entities.UserParams{TeamName: teamData.TeamName},
	)
	if err != nil {
		return
	}

	res.Id = selectedTeam.Id
	res.Name = selectedTeam.Name
	res.Users = selectedUsers

	return
}
