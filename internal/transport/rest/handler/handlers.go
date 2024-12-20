package handler

import (
	"net/http"

	messenger_service "github.com/KrizzMU/coolback-alkol/internal/core/messenger/domain/service"

	"github.com/gorilla/websocket"

	"github.com/KrizzMU/coolback-alkol/internal/service"
	"github.com/KrizzMU/coolback-alkol/pkg/auth"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/KrizzMU/coolback-alkol/docs"
)

// Handler
// @title Messenger API
// @version 0.1
// @BasePath /api/v1
// @securityDefinitions.apikey  BearerAuth
// @in              header
// @name            Authorization
// @description     "Укажите 'Bearer', а затем ваш JWT токен."
type Handler struct {
	tokenManger      auth.TokenManager
	services         *service.Service
	upgrader         *websocket.Upgrader
	messengerService *messenger_service.Messenger
}

func New(
	tokenManager auth.TokenManager,
	services *service.Service,
	messengerService *messenger_service.Messenger,
) *Handler {
	return &Handler{
		tokenManger:      tokenManager,
		services:         services,
		messengerService: messengerService,
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
		chat.Handle(http.MethodPost, "/add/members", h.AddMember)
		chat.Handle(http.MethodDelete, "/:id", h.DeleteChat)
		chat.Handle(http.MethodGet, "/messages", h.GetChatMessages)
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
		user.Handle(http.MethodPost, "/set/username", h.isLogedIn, h.SetUsername)
	}

	messenger := v1.Group("/messenger", h.isLogedIn)
	{
		messenger.Handle(http.MethodGet, "/connect", h.Connect)
	}

	return r
}
