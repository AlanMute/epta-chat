package handler

import (
	"github.com/KrizzMU/coolback-alkol/internal/core/messenger/domain/model"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/KrizzMU/coolback-alkol/docs"
)

// @title Messenger API
// @version 0.1
// @BasePath /api/v1

type Handler struct {
	messenger *model.Messenger
}

func New(messenger *model.Messenger) *Handler {
	return &Handler{
		messenger: messenger,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	// Middlewares
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("api/v1")

	messengerGroup := v1.Group("/messenger")

	messengerGroup.GET("/connect", h.Connect)

	return r
}
