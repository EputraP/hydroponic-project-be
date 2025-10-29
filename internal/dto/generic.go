package dto

import "github.com/google/uuid"

type OnlyIdResponse struct {
	ID uuid.UUID `json:"id"`
}

type IdAndNameResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type IdNameAndDescriptionResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type PaginationRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}
