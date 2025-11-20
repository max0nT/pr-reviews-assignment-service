package teammanage

import (
	"github.com/max0nT/pr-assign/internal/repo"

	"github.com/max0nT/pr-assign/pkg/postgres"
)

type TeamManage struct {
	Cfg      *postgres.Postgres
	UserRepo repo.UserRepository
	TeamRepo repo.TeamRepository
}

func New(
	cfg *postgres.Postgres,
	ur repo.UserRepository,
	tr repo.TeamRepository,
) *TeamManage {
	return &TeamManage{
		Cfg:      cfg,
		UserRepo: ur,
		TeamRepo: tr,
	}
}
