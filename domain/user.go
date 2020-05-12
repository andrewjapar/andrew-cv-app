package domain

import (
	"time"
)

type User struct {
	ID                   int64  `json:"id"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	Email                string `gorm:"unique_index" json:"email"`
	Password             string `json:"password"`
	RegistrationPlatform string `json:"registration_platform"`
	Wedding              []*Wedding
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type UserRepository interface {
	Get() ([]User, error)
}
