package repo

import "github.com/max0nT/pr-assign/internal/entities"

type (
	TeamRepository interface {
		InsertTeam(string) (entities.ItemSimple, error)
	}

	UserRepository interface {
		InsertUsers([]entities.User, string) ([]entities.User, error)
	}
)
