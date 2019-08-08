package models

import (
	"time"
)

// Charge : charge object
type Charge struct {
	ID            int       `json:"id" db:"id"`
	ExternalID    string    `json:"external_id" db:"external_id"`
	AmountInCents int64     `json:"amount_in_cents" db:"amount_in_cents"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
