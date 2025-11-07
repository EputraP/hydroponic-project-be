package dto

import "github.com/google/uuid"

type PlantGrowth struct {
	TowerId           uuid.UUID `json:"tower_id"`
	ProcessId         uuid.UUID `json:"process_id"`
	Cycle             int       `json:"cycle"`
	HSS               int       `json:"hss"`
	HST               int       `json:"hst"`
	Height            int       `json:"height"`
	PH                float32   `json:"ph"`
	PPM               float32   `json:"ppm"`
	WaterLevel        int       `json:"water_level"`
	OVRPlantCondition string    `json:"ovr_plant_condition"`
	Remarks           string    `json:"remarks"`
	PlantAmount       int       `json:"plant_amount"`
}
