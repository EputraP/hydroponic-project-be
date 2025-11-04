package dto

import "github.com/google/uuid"

type Asset struct {
	UOMID       uuid.UUID `json:"uom_id"`
	PlantID     uuid.UUID `json:"plant_id"`
	AssetTypeID uuid.UUID `json:"asset_type_id"`
	AssetName   string    `json:"asset_name"`
	Price       int       `json:"price"`
	Value       int       `json:"value"`
	Cycle       int       `json:"cycle"`
}
