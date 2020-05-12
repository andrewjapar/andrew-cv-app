package domain

import "time"

type WeddingOrganizer struct {
	ID        int64 `json:"id"`
	User      *User
	UserID    int64
	Wedding   *Wedding
	WeddingID int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
