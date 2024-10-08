package repository

import (
	"time"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type Contact interface {
	Add(ownerId, contactId uint64) error
	Delete(ownerId, contactId uint64) error
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
	SignIn(user core.User) (uint64, error)
	SignUp(user core.User) error
	AddSession(session core.Session) error
	CheckRefresh(token string) error
}

type Message interface {
	Send(text string, senderId, chatId uint64, sendingTime time.Time) error
}

type Repository struct {
	Contact
	Chat
	User
	Message
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Contact: NewContactPostgres(db),
		Chat:    NewChatPostgres(db),
		User:    NewUserPostgres(db),
		Message: NewMessagePostgres(db),
	}
}
