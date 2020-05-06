package domain

import (
	"context"
	"time"
)

type Profile struct {
	ID          int64     `json:"id"`
	Avatar      string    `json:"avatar" validate:"required"`
	FullName    string    `json:"full_name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Quote       string    `json:"quote" validate:"required"`
	GithubURL   string    `json:"github_url" validate:"required"`
	TwitterURL  string    `json:"twitter_url" validate:"required"`
	LinkedinURL string    `json:"linkedin_url" validate:"required"`
	MediumURL   string    `json:"medium_url" validate:"required"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProfileRepository interface {
	Get(ctx context.Context) (Profile, error)
}
