package setup

import (
	handler "hydroponic-be/internal/handler/growing"
	"hydroponic-be/internal/routes"
	service "hydroponic-be/internal/service/growing"
)

func prepareGrowing(db *DBs) (handlers routes.HandlersGrowing) {

	assetRepo := *db.Admin.AssetRepo
	processRepo := *db.Admin.ProcessRepo
	plantGrowthRepo := *db.Growing.PlantGrowthRepo
	unhealthyPlantTreatmentRepo := *db.Growing.UnhealthyPlantTreatmentRepo

	plantService := service.NewPlantGrowthService(service.PlantGrowthServiceConfig{
		PlantGrowthRepo: plantGrowthRepo,
		AssetRepo:       assetRepo,
		ProcessRepo:     processRepo,
	})
	unhealthyPlantTreatmentService := service.NewUnhealthyPlantTreatmentService(service.UnhealthyPlantTreatmentServiceConfig{
		PlantGrowthRepo: unhealthyPlantTreatmentRepo,
	})

	plantGrowthHandler := handler.NewPlantGrowthHandler(handler.PlantGrowthHandlerConfig{
		PlantGrowthService: plantService,
	})
	unhealthyPlantTreatmentHandler := handler.NewUnhealthyPlantTreatmentHandler(handler.UnhealthyPlantTreatmentHandlerConfig{
		UnhealthyPlantTreatmentService: unhealthyPlantTreatmentService,
	})

	handlers = routes.HandlersGrowing{
		PlantGrowth:             plantGrowthHandler,
		UnhealthyPlantTreatment: unhealthyPlantTreatmentHandler,
	}

	return
}
