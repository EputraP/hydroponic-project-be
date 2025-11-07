package model

import "github.com/google/uuid"

type Remark struct {
	ID          uuid.UUID `json:"id"`
	Remark      string    `json:"remark"`
	Description string    `json:"description"`
}
