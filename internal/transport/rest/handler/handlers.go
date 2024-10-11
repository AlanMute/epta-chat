package handler

import (
	"net/http"

	"github.com/KrizzMU/coolback-alkol/internal/core/messenger/domain/model"
	"github.com/gorilla/websocket"

	"github.com/KrizzMU/coolback-alkol/internal/service"
	"github.com/KrizzMU/coolback-alkol/pkg/auth"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/KrizzMU/coolback-alkol/docs"
)

// @title Messenger API
// @version 0.1
// @BasePath /api/v1

type Handler struct {
	tokenManger auth.TokenManager
	services    *service.Service
	upgrader    *websocket.Upgrader
	messenger   *model.Messenger
}

func New(
	tokenManager auth.TokenManager,
	services *service.Service,
	messenger *model.Messenger,
) *Handler {
	return &Handler{
		tokenManger: tokenManager,
		services:    services,
		messenger:   messenger,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	// Middlewares
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("api/v1")

	chat := v1.Group("/chat", h.isLogedIn)
	{
		chat.Handle(http.MethodGet, "/all", h.GetChats)
		chat.Handle(http.MethodGet, "/:id", h.GetChatById)
		chat.Handle(http.MethodGet, "/members/:id", h.GetChatMembers)
		chat.Handle(http.MethodPost, "/", h.AddChat)
		chat.Handle(http.MethodDelete, "/:id", h.DeleteChat)
	}

	contact := v1.Group("/contact", h.isLogedIn)
	{
		contact.Handle(http.MethodGet, "/all", h.GetContacts)
		contact.Handle(http.MethodGet, "/:id", h.GetContactById)
		contact.Handle(http.MethodPost, "/", h.AddContact)
		contact.Handle(http.MethodDelete, "/:id", h.DeleteContact)
	}

	user := v1.Group("/user")
	{
		user.Handle(http.MethodPost, "/sign-in", h.SignIn)
		user.Handle(http.MethodPost, "/sign-up", h.SignUp)
		user.Handle(http.MethodPost, "/refresh", h.Refresh)
	}

	messenger := v1.Group("/messenger", h.isLogedIn)
	{
		messenger.Handle(http.MethodGet, "/connect", h.Connect)
	}

	return r
}
