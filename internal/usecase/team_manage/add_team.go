package teammanage

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/max0nT/pr-assign/internal/entities"
)

func (tm *TeamManage) AddTeam(
	teamData entities.ItemCreate,
) (res entities.ItemRead, err error) {
	insertedTeam, err := tm.TeamRepo.InsertTeam(teamData.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = &entities.RequestError{
				Msg: fmt.Sprintf(
					"Team with %s already exists",
					insertedTeam.Name,
				),
				StatusCode: http.StatusBadRequest,
			}
		}
		return
	}

	insertedUsers, err := tm.UserRepo.InsertUsers(
		teamData.Users,
		insertedTeam.Name,
	)
	if err != nil {
		return
	}

	res.Id = insertedTeam.Id
	res.Name = insertedTeam.Name
	res.Users = insertedUsers

	return
}
