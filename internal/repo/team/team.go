package teamrepo

import "github.com/max0nT/pr-assign/pkg/postgres"

type TeamRepository struct {
	Cfg *postgres.Postgres
}

func New(cfg *postgres.Postgres) *TeamRepository {
	return &TeamRepository{
		Cfg: cfg,
	}
}
