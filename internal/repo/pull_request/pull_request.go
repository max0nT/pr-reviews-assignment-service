package prrepo

import "github.com/max0nT/pr-assign/pkg/postgres"

type PrRepository struct {
	Cfg *postgres.Postgres
}

func New(cfg *postgres.Postgres) *PrRepository {
	return &PrRepository{
		Cfg: cfg,
	}
}
