package prmanage

import (
	"github.com/max0nT/pr-assign/internal/repo"

	"github.com/max0nT/pr-assign/pkg/postgres"
)

type PrManage struct {
	Cfg      *postgres.Postgres
	UserRepo repo.UserRepository
	TeamRepo repo.TeamRepository
	PrRepo   repo.PrRepository
}

func New(
	cfg *postgres.Postgres,
	ur repo.UserRepository,
	tr repo.TeamRepository,
	pr repo.PrRepository,
) *PrManage {
	return &PrManage{
		Cfg:      cfg,
		UserRepo: ur,
		TeamRepo: tr,
		PrRepo:   pr,
	}
}
