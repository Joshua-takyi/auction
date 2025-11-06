package models

import "github.com/google/uuid"

type Bidder struct {
	ID uuid.UUID `db:"id" json:"id"`
}
