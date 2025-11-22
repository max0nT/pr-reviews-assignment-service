package entities

type User struct {
	Id       string `json:"id"        validate:"required,min=3,max=50"`
	Username string `json:"username"  validate:"required,min=3,max=50"`
	TeamName string `json:"team_name"`
	IsActive bool   `json:"is_active" validate:"required"`
}

type UserStats struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	TeamName string `json:"team_name"`
	IsActive bool   `json:"is_active"`
	PrCount  int    `json:"pr_count"`
	RwCount  int    `json:"rw_count"`
}

type UserParams struct {
	Id       string   `form:"id"`
	NotId    string   `form:"not_id"`
	IdIn     []string `form:"id_in"`
	NotIdIn  []string `form:"not_id_in"`
	Username string   `form:"username"`
	TeamName string   `form:"team_name"`
	IsActive bool     `form:"is_active"`
	Limit    int      `form:"limit"`
}

type UserChangeActive struct {
	Id       string `json:"id"`
	IsActive bool   `json:"is_active"`
}
