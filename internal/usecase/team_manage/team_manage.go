package teammanage

import "github.com/max0nT/pr-assign/internal/repo"

type TeamManage struct {
	UserRepo repo.UserRepository
	TeamRepo repo.TeamRepository
}

func New(ur repo.UserRepository, tr repo.TeamRepository) *TeamManage {
	return &TeamManage{
		UserRepo: ur,
		TeamRepo: tr,
	}
}
