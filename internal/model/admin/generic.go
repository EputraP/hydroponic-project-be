package model

import "github.com/google/uuid"

type OnlyIdResponse struct {
	ID uuid.UUID `json:"id"`
}
