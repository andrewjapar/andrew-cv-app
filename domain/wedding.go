package domain

import (
	"time"
)

type Wedding struct {
	ID          int64     `json:"id"`
	Title       string    `json:"email"`
	Description string    `json:"description"`
	WeddingDate time.Time `json:"wedding_date"`
	Code        string    `gorm:"unique_index;not null" json:"code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WeddingRepository interface {
	GetByUserID(id int64) ([]Wedding, error)
}
