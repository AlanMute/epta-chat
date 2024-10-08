package repository

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type Contact interface {
	Add(ownerId, conactId uint64) error
	Delete(ownerId, conactId uint64) error
	GetAll(ownerId uint64) ([]core.UserInfo, error)
	GetById(id uint64) (core.UserInfo, error)
}

type Chat interface {
	Add(name string, isDirect bool, ownerId uint64, members []uint64) error
	Delete(userId, chatId uint64) error
	GetById(userId, chatId uint64) (core.Chat, error)
	GetAll(userId uint64) ([]core.Chat, error)
	GetMembers(userId, chatId uint64) ([]core.UserInfo, error)
}

type User interface {
	Add(core.User) error
}

type Session interface {
	Add(session core.Session) error
	CheckRefresh(token string) error
}

type Repository struct {
	Contact
	Chat
	User
	Session
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		//Contact: NewContactPostgres(db),
		//Chat:    NewChatPostgres(db),
		//User:    NewUserPostgres(db),
		//Session: NewLessonPostgres(db),
	}
}
