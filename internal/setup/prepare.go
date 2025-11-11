package setup

import (
	"hydroponic-be/internal/routes"
	dbstore "hydroponic-be/internal/store"
	"hydroponic-be/internal/util/logger"
)

func Prepare() (handlers routes.Handlers, middlewares routes.Middlewares) {
	logger.Info("main", "Initializing dependencies...", nil)
	db := dbstore.Get()

	dbs := prepareDb(db)

	handlersAdmin := prepareAdmin(dbs)
	handlersGrowing := prepareGrowing(dbs)
	handlerTransaction := prepareTransaction(dbs)

	handlers = routes.Handlers{
		Admin:       handlersAdmin,
		Growing:     handlersGrowing,
		Transaction: handlerTransaction,
	}

	logger.Info("main", "Application initialized successfully.", nil)
	return
}
