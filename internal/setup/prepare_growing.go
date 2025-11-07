package setup

import (
	handler "hydroponic-be/internal/handler/growing"
	"hydroponic-be/internal/routes"
	service "hydroponic-be/internal/service/growing"
)

func prepareGrowing(db *DBs) (handlers routes.HandlersGrowing) {

	plantGrowthRepo := *db.Growing.PlantGrowthRepo
	assetRepo := *db.Admin.AssetRepo
	processRepo := *db.Admin.ProcessRepo

	plantService := service.NewPlantGrowthService(service.PlantGrowthServiceConfig{
		PlantGrowthRepo: plantGrowthRepo,
		AssetRepo:       assetRepo,
		ProcessRepo:     processRepo,
	})

	plantGrowthHandler := handler.NewPlantGrowthHandler(handler.PlantGrowthHandlerConfig{
		PlantGrowthService: plantService,
	})

	handlers = routes.HandlersGrowing{
		PlantGrowth: plantGrowthHandler,
	}

	return
}
