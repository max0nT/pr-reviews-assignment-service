package entities

type TeamCreate struct {
	Name  string `json:"name"  validate:"required,min=3,max=20"`
	Users []User `json:"users" validate:"required,min=1,max=20"`
}

type TeamRead struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"members"`
}

type TeamSimple struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
