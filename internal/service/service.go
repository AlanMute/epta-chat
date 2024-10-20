package service

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"github.com/KrizzMU/coolback-alkol/pkg/auth"
)

type Contact interface {
	Add(ownerId uint64, contactLogin string) error
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
	AddMember(ownerId, chatId uint64, members []uint64) error
}

type User interface {
	SignIn(login, password string) (uint64, core.Tokens, error)
	SignUp(login, password string) error
	Refresh(userId uint64, refreshToken string) (string, error)
	SetUserName(userId uint64, userName string) error
}

type Messgae interface {
	GetBatch(userId, chatId, pageNumber uint64) ([]core.Message, error)
}

type Service struct {
	Contact
	Chat
	User
	Messgae
}

func New(repo *repository.Repository, t auth.TokenManager) *Service {
	return &Service{
		Contact: NewContactService(repo.Contact),
		Chat:    NewChatService(repo.Chat),
		User:    NewUserService(repo.User, t),
		Messgae: NewMessageService(repo.Message),
	}
}
