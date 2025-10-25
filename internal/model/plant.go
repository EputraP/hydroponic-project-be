package model

import "github.com/google/uuid"

type Plant struct {
	ID              uuid.UUID `json:"id"`
	PlantName       string    `json:"plant_name"`
	ScientificName  string    `json:"scientific_name"`
	Variety         string    `json:"variety"`
	PlantType       string    `json:"plant_type"`
	Description     string    `json:"description"`
	PHMin           float64   `json:"pH_min"`
	PHMax           float64   `json:"pH_max"`
	PPMMin          int       `json:"ppm_min"`
	PPMMax          int       `json:"ppm_max"`
	LightHours      float64   `json:"light_hours"`
	TemperatureMin  float64   `json:"temperature_min"`
	TemperatureMax  float64   `json:"temperature_max"`
	HarvestDays     int       `json:"harvest_days"`
	GerminationDays int       `json:"germination_days"`
	HSSDays         int       `json:"hss_days"`
	HSTDays         int       `json:"hst_days"`
}
