package setup

import (
	handler "hydroponic-be/internal/handler/growing"
	repository "hydroponic-be/internal/repository/growing"
	"hydroponic-be/internal/routes"
	service "hydroponic-be/internal/service/growing"

	"gorm.io/gorm"
)

func prepareGrowing(db *gorm.DB) (handlers routes.HandlersGrowing) {

	plantGrowthRepo := repository.NewPlantGrowthRepository(db)

	plantService := service.NewPlantGrowthService(service.PlantGrowthServiceConfig{
		PlantGrowthRepo: plantGrowthRepo,
	})

	plantGrowthHandler := handler.NewPlantGrowthHandler(handler.PlantGrowthHandlerConfig{
		PlantGrowthService: plantService,
	})

	handlers = routes.HandlersGrowing{
		PlantGrowth: plantGrowthHandler,
	}

	return
}
