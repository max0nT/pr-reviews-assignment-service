package teamrepo

import "github.com/max0nT/pr-assign/pkg/postgres"

type ItemRepository struct {
	Cfg *postgres.Postgres
}

func New(cfg *postgres.Postgres) *ItemRepository {
	return &ItemRepository{
		Cfg: cfg,
	}
}
