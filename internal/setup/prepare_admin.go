package setup

import (
	handler "hydroponic-be/internal/handler/admin"
	"hydroponic-be/internal/routes"
	service "hydroponic-be/internal/service/admin"
)

func prepareAdmin(db *DBs) (handlers routes.HandlersAdmin) {

	PlantRepo := *db.Admin.PlantRepo
	processRepo := *db.Admin.ProcessRepo
	remarkRepo := *db.Admin.RemarkRepo
	uomRepo := *db.Admin.UomRepo
	assetTypeRepo := *db.Admin.AssetTypeRepo
	assetRepo := *db.Admin.AssetRepo

	plantService := service.NewPlantService(service.PlantServiceConfig{
		PlantRepo: PlantRepo,
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
		PlantRepo: PlantRepo,
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
