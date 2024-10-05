package service

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"github.com/KrizzMU/coolback-alkol/pkg/auth"
)

type Contact interface {
	Add(ownerId, conactId uint64) error
	Delete(ownerId, conactId uint64) error
	GetAll(ownerId uint64) ([]core.UserInfo, error)
	GetById(id uint64) (core.UserInfo, error)
}

type Chat interface {
	Add(name string, isDirect bool, ownerId uint64, members []uint64) (uint64, error)
	Delete(userId, chatId uint64) error
	GetById(userId, chatId uint64) (core.Chat, error)
	GetAll(userId uint64) ([]core.Chat, error)
	GetMembers(userId, chatId uint64) ([]core.UserInfo, error)
}

type User interface {
	Add(login, password string) error
}

type Session interface {
	Add(session core.Session) error
	CheckRefresh(token string) error
}

type Service struct {
	Contact
	Chat
	User
	Session
}

func NewRepository(repo *repository.Repository, t auth.TokenManager) *Service {
	return &Service{
		Contact: NewContactSevice(repo.Contact),
		Chat:    NewChatSevice(repo.Chat),
		User:    NewUserSevice(repo.User),
		Session: NewSessionService(repo.Session),
	}
}
