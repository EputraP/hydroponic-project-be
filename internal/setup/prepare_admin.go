package setup

import (
	handler "hydroponic-be/internal/handler/admin"
	repository "hydroponic-be/internal/repository/admin"
	"hydroponic-be/internal/routes"
	service "hydroponic-be/internal/service/admin"

	"gorm.io/gorm"
)

func prepareAdmin(db *gorm.DB) (handlers routes.HandlersAdmin) {

	plantRepo := repository.NewPlantRepository(db)
	processRepo := repository.NewProcessRepository(db)
	remarkRepo := repository.NewRemarkRepository(db)
	uomRepo := repository.NewUomRepository(db)
	assetTypeRepo := repository.NewAssetTypeRepository(db)
	assetRepo := repository.NewAssetRepository(db)

	plantService := service.NewPlantService(service.PlantServiceConfig{
		PlantRepo: plantRepo,
	})
	processService := service.NewProcessService(service.ProcessServiceConfig{
		ProcessRepo: processRepo,
	})
	remarkService := service.NewRemarkService(service.RemarkServiceConfig{
		RemarkRepo: remarkRepo,
	})
	uomService := service.NewUomService(service.UomServiceConfig{
		UomRepo: uomRepo,
	})
	assetTypeService := service.NewAssetTypeService(service.AssetTypeServiceConfig{
		AssetTypeRepo: assetTypeRepo,
	})
	assetService := service.NewAssetService(service.AssetServiceConfig{
		AssetRepo: assetRepo,
		PlantRepo: plantRepo,
		UomRepo:   uomRepo,
		AssetType: assetTypeRepo,
	})

	plantHandler := handler.NewPlantHandler(handler.PlantHandlerConfig{
		PlantService: plantService,
	})
	processHandler := handler.NewProcessHandler(handler.ProcessHandlerConfig{
		ProcessService: processService,
	})
	remarkHandler := handler.NewRemarkHandler(handler.RemarkHandlerConfig{
		RemarkService: remarkService,
	})
	uomHandler := handler.NewUomHandler(handler.UomHandlerConfig{
		UomService: uomService,
	})
	assetTypeHandler := handler.NewAssetTypeHandler(handler.AssetTypeHandlerConfig{
		AssetTypeService: assetTypeService,
	})
	assetHandler := handler.NewAssetHandler(handler.AssetHandlerConfig{
		AssetService: assetService,
	})

	handlers = routes.HandlersAdmin{
		Plant:     plantHandler,
		Process:   processHandler,
		Remark:    remarkHandler,
		Uom:       uomHandler,
		AssetType: assetTypeHandler,
		Asset:     assetHandler,
	}

	return
}
