package setup

import (
	handler "hydroponic-be/internal/handler/growing"
	"hydroponic-be/internal/routes"
	service "hydroponic-be/internal/service/growing"
)

func prepareGrowing(db *DBs) (handlers routes.HandlersGrowing) {

	plantGrowthRepo := *db.Growing.PlantGrowthRepo
	assetRepo := *db.Admin.AssetRepo

	plantService := service.NewPlantGrowthService(service.PlantGrowthServiceConfig{
		PlantGrowthRepo: plantGrowthRepo,
		AssetRepo:       assetRepo,
	})

	plantGrowthHandler := handler.NewPlantGrowthHandler(handler.PlantGrowthHandlerConfig{
		PlantGrowthService: plantService,
	})

	handlers = routes.HandlersGrowing{
		PlantGrowth: plantGrowthHandler,
	}

	return
}
