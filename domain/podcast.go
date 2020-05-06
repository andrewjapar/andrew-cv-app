package domain

import (
	"context"
	"time"
)

type Podcast struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	AudioURL    string    `json:"audio_url" validate:"required"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type PodcastRepository interface {
	Get(ctx context.Context) ([]Podcast, error)
}
