package model

import "github.com/google/uuid"

type Process struct {
	ID          uuid.UUID `json:"id"`
	ProcessName string    `json:"process_name"`
	Description string    `json:"description"`
}
