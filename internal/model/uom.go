package model

import "github.com/google/uuid"

type Uom struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Symbol      string    `json:"symbol"`
	Description string    `json:"description"`
}
