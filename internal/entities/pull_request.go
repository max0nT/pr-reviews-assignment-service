package entities

import "time"

type PrCreate struct {
	PrId      string `json:"pull_request_id"   validate:"required,min=3,max=20"`
	PrName    string `json:"pull_request_name" validate:"required,min=3,max=20"`
	CreatedBy string `json:"created_by_id"     validate:"required,min=3,max=20"`
}

type PrSimple struct {
	PrId      string    `json:"pull_request_id"`
	PrName    string    `json:"pull_request_name"`
	CreatedBy string    `json:"created_by_id"`
	IsMerged  bool      `json:"is_merged"`
	CreatedAt time.Time `json:"created_at"`
	MergedAt  time.Time `json:"merged_at"`
}

type PrRead struct {
	PrId      string    `json:"pull_request_id"`
	PrName    string    `json:"pull_request_name"`
	CreatedBy string    `json:"created_by_id"`
	IsMerged  bool      `json:"is_merged"`
	CreatedAt time.Time `json:"created_at"`
	MergedAt  time.Time `json:"merged_at"`
	Reviewers []User    `json:"reviewers"`
}

type PrMerge struct {
	PrId string `json:"pull_request_id" validate:"required"`
}
