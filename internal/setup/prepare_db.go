package setup

import (
	repositoryAdmin "hydroponic-be/internal/repository/admin"
	repositoryGrowing "hydroponic-be/internal/repository/growing"

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
	PlantGrowthRepo *repositoryGrowing.PlantGrowthRepository
}
type DBs struct {
	Admin   *DBsAdmin
	Growing *DBsGrowing
}

func prepareDb(db *gorm.DB) (dbs *DBs) {

	PlantRepo := repositoryAdmin.NewPlantRepository(db)
	processRepo := repositoryAdmin.NewProcessRepository(db)
	remarkRepo := repositoryAdmin.NewRemarkRepository(db)
	uomRepo := repositoryAdmin.NewUomRepository(db)
	assetTypeRepo := repositoryAdmin.NewAssetTypeRepository(db)
	assetRepo := repositoryAdmin.NewAssetRepository(db)

	plantGrowthRepo := repositoryGrowing.NewPlantGrowthRepository(db)

	dbsAdmin := &DBsAdmin{
		PlantRepo:     &PlantRepo,
		ProcessRepo:   &processRepo,
		RemarkRepo:    &remarkRepo,
		UomRepo:       &uomRepo,
		AssetTypeRepo: &assetTypeRepo,
		AssetRepo:     &assetRepo,
	}

	dbsGrowing := &DBsGrowing{
		PlantGrowthRepo: &plantGrowthRepo,
	}

	dbs = &DBs{
		Admin:   dbsAdmin,
		Growing: dbsGrowing,
	}

	return
}
