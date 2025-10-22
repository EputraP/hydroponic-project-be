package routes

import "github.com/gin-gonic/gin"

type Handlers struct {
	// Define your route handlers here
}

type Middlewares struct {
	// Define your middlewares here
}

func Build(srv *gin.Engine, handlers Handlers, middlewares Middlewares) {

}
