package model

import "github.com/google/uuid"

type Asset struct {
	ID          uuid.UUID `json:"id"`
	UOMID       uuid.UUID `json:"uom_id"`
	PlantID     uuid.UUID `json:"plant_id"`
	AssetTypeID uuid.UUID `json:"asset_type_id"`
	AssetName   string    `json:"asset_name"`
	Price       float64   `json:"price"`
	Value       float64   `json:"value"`
	Cycle       int       `json:"cycle"`
	CreatedAt   string    `json:"created_at"`
}
