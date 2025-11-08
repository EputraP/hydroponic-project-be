package dto

import "github.com/google/uuid"

type UnhealthyPlantTreatment struct {
	TowerId        uuid.UUID `json:"tower_id"`
	Cycle          int       `json:"cycle"`
	TowerHoleNo    int       `json:"tower_hole_no"`
	Treatment      string    `json:"treatment"`
	AfterTreatment string    `json:"after_treatment"`
}
