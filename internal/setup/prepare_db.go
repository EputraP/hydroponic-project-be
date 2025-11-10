package setup

import (
	repositoryAdmin "hydroponic-be/internal/repository/admin"
	repositoryGrowing "hydroponic-be/internal/repository/growing"
	repositoryTransaction "hydroponic-be/internal/repository/transaction"

	"gorm.io/gorm"
)

type DBsAdmin struct {
	PlantRepo     *repositoryAdmin.PlantRepository
	ProcessRepo   *repositoryAdmin.ProcessRepository
	RemarkRepo    *repositoryAdmin.RemarkRepository
	UomRepo       *repositoryAdmin.UomRepository
	AssetTypeRepo *repositoryAdmin.AssetTypeRepository
	AssetRepo     *repositoryAdmin.AssetRepository
}
type DBsGrowing struct {
	PlantGrowthRepo             *repositoryGrowing.PlantGrowthRepository
	UnhealthyPlantTreatmentRepo *repositoryGrowing.UnhealthyPlantTreatmentRepository
}
type DBsTransaction struct {
	AssetOpsTransactionRepo *repositoryTransaction.AssetOpsTransactionRepository
}
type DBs struct {
	Admin       *DBsAdmin
	Growing     *DBsGrowing
	Transaction *DBsTransaction
}

func prepareDb(db *gorm.DB) (dbs *DBs) {

	PlantRepo := repositoryAdmin.NewPlantRepository(db)
	processRepo := repositoryAdmin.NewProcessRepository(db)
	remarkRepo := repositoryAdmin.NewRemarkRepository(db)
	uomRepo := repositoryAdmin.NewUomRepository(db)
	assetTypeRepo := repositoryAdmin.NewAssetTypeRepository(db)
	assetRepo := repositoryAdmin.NewAssetRepository(db)

	plantGrowthRepo := repositoryGrowing.NewPlantGrowthRepository(db)
	unhealthyPlantTreatmentRepo := repositoryGrowing.NewUnhealthyPlantTreatmentRepository(db)

	assetOpsTransactionRepo := repositoryTransaction.NewAssetOpsTransactionRepository(db)

	dbsAdmin := &DBsAdmin{
		PlantRepo:     &PlantRepo,
		ProcessRepo:   &processRepo,
		RemarkRepo:    &remarkRepo,
		UomRepo:       &uomRepo,
		AssetTypeRepo: &assetTypeRepo,
		AssetRepo:     &assetRepo,
	}

	dbsGrowing := &DBsGrowing{
		PlantGrowthRepo:             &plantGrowthRepo,
		UnhealthyPlantTreatmentRepo: &unhealthyPlantTreatmentRepo,
	}

	dbsTransaction := &DBsTransaction{
		AssetOpsTransactionRepo: &assetOpsTransactionRepo,
	}

	dbs = &DBs{
		Admin:       dbsAdmin,
		Growing:     dbsGrowing,
		Transaction: dbsTransaction,
	}

	return
}
