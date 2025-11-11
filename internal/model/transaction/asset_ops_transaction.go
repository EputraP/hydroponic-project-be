package model

import "github.com/google/uuid"

type AssetOpsTransaction struct {
	ID      uuid.UUID `json:"id"`
	TowerId uuid.UUID `json:"tower_id"`
	AssetId uuid.UUID `json:"asset_id"`
	Cycle   int       `json:"cycle"`
	Value   int       `json:"value"`
}
