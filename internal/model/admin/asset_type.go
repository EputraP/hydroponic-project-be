package model

import "github.com/google/uuid"

type AssetType struct {
	ID          uuid.UUID `json:"id"`
	TypeName    string    `json:"type_name"`
	Description string    `json:"description"`
}
