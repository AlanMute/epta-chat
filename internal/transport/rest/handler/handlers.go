package handler

import (
	"net/http"

	"github.com/KrizzMU/coolback-alkol/internal/service"
	"github.com/KrizzMU/coolback-alkol/pkg/auth"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Messenger API
// @version 0.1
// @BasePath /api/v1

type Handler struct {
	tokenManger auth.TokenManager
	services    *service.Service
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	// Middlewares
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	chat := r.Group("/chat") //TODO: make middleware
	{
		chat.Handle(http.MethodGet, "/all", h.GetChats)
		chat.Handle(http.MethodGet, "/:id", h.GetChatById)
		chat.Handle(http.MethodGet, "/members/:id", h.GetChatMembers)
		chat.Handle(http.MethodPost, "/", h.AddChat)
		chat.Handle(http.MethodDelete, "/:id", h.DeleteChat)
	}

	contacts := r.Group("/contacts")
	{
		contacts.Handle(http.MethodGet, "/all", h.GetContacts)
		contacts.Handle(http.MethodGet, "/:id", h.GetContactById)
		chat.Handle(http.MethodPost, "/", h.AddContact)
		chat.Handle(http.MethodDelete, "/:id", h.DeleteContact)
	}

	user := r.Group("/user")
	{
		user.Handle(http.MethodPost, "/sign-in", h.SignIn)
		user.Handle(http.MethodPost, "/sign-up", h.SignUp)
		user.Handle(http.MethodPost, "/refresh", h.Refresh)
	}

	return r
}
