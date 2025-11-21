package entities

type User struct {
	Id       string `json:"id"        validate:"required,min=3,max=5"`
	Username string `json:"username"  validate:"required,min=3,max=20"`
	TeamName string `json:"team_name"`
	IsActive bool   `json:"is_active" validate:"required"`
}

type UserParams struct {
	Id       string   `json:"id"`
	NotId    string   `json:"not_id"`
	IdIn     []string `json:"id_in"`
	Username string   `json:"username"`
	TeamName string   `json:"team_name"`
	IsActive bool     `json:"is_active"`
	Limit    int      `json:"limit"`
}
