package userrepo

import (
	"github.com/max0nT/pr-assign/pkg/postgres"
)

type UserRepository struct {
	Cfg *postgres.Postgres
}

func New(cfg *postgres.Postgres) *UserRepository {
	return &UserRepository{
		Cfg: cfg,
	}
}
