package models

import (
	"time"
)

type CountryStatus struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	CountryISO string    `json:"country_iso"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}