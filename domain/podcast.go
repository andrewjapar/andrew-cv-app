package domain

import (
	"time"
)

type Podcast struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	AudioURL    string    `json:"audio_url" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PodcastRepository interface {
	Get() ([]Podcast, error)
}
