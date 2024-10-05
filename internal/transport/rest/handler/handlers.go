package handler

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/KrizzMU/coolback-alkol/docs"
)

// @title Messenger API
// @version 0.1
// @BasePath /api/v1

type Handler struct {
	// TODO: Add dependencies
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	// Middlewares
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
