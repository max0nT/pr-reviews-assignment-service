package entities

import "time"

type PrCreate struct {
	PrId      string `json:"pull_request_id"   validate:"required,min=3,max=20"`
	PrName    string `json:"pull_request_name" validate:"required,min=3,max=20"`
	CreatedBy string `json:"created_by_id"     validate:"required,min=3,max=20"`
}

type PrSimple struct {
	PrId          string     `json:"pull_request_id"`
	PrName        string     `json:"pull_request_name"`
	CreatedBy     string     `json:"created_by_id"`
	IsMerged      bool       `json:"is_merged"`
	CreatedAt     time.Time  `json:"created_at"`
	MergedAt      *time.Time `json:"merged_at"`
	CreatedByData User       `json:"created_by"`
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

type PrParams struct {
	PrId         string
	CreatedBy    string
	ReviewerIdIn []string
	IsMerged     bool
}

type PrReviewerParams struct {
	PrId string
}

type PrReviewer struct {
	PrId       string
	ReviewerId string
}

type PrMerge struct {
	PrId string `json:"pull_request_id" validate:"required"`
}

type PrUnassign struct {
	PrId      string `json:"pull_request_id" validate:"required"`
	OldUserId string `json:"old_reviewer_id" validate:"required"`
}

type PrAssign struct {
	PrId      string `json:"pull_request_id" validate:"required"`
	NewUserId string `json:"new_reviewer_id" validate:"required"`
}
