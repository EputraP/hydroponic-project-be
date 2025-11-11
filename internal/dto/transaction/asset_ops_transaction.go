package dto

import "github.com/google/uuid"

type AssetOpsTransaction struct {
	TowerId uuid.UUID `json:"tower_id"`
	AssetId uuid.UUID `json:"asset_id"`
	Cycle   int       `json:"cycle"`
	Value   int       `json:"value"`
}
