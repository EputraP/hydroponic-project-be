package setup

import (
	"hydroponic-be/internal/routes"
	dbstore "hydroponic-be/internal/store"
	"hydroponic-be/internal/util/logger"
)

func Prepare() (handlersAdmin routes.HandlersAdmin, middlewares routes.Middlewares) {
	logger.Info("main", "Initializing dependencies...", nil)
	db := dbstore.Get()

	handlersAdmin = prepareAdmin(db)
	logger.Info("main", "Application initialized successfully.", nil)
	return
}
