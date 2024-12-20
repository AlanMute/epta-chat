package repository

import (
	"time"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type Contact interface {
	Add(ownerId uint64, contactLogin string) error
	Delete(ownerId, contactId uint64) error
	GetAll(ownerId uint64) ([]core.UserInfo, error)
	GetById(id uint64) (core.UserInfo, error)
}

type Chat interface {
	Add(name string, isDirect bool, ownerId uint64, members []uint64) (uint64, error)
	AddMember(ownerId, chatId uint64, members []uint64) error
	Delete(userId, chatId uint64) error
	GetById(userId, chatId uint64) (core.Chat, error)
	GetAll(userId uint64) ([]core.Chat, error)
	GetMembers(userId, chatId uint64) ([]core.UserInfo, error)
	EnsureCommonChatExists() error
	FetchAllChatIDs() ([]uint64, error)
}

type User interface {
	SignIn(user core.User) (uint64, error)
	SignUp(user core.User) error
	SetUserName(userId uint64, userName string) error
	AddSession(session core.Session) error
	CheckRefresh(token string) error
	GetById(userId uint64) (core.User, error)
}

type Message interface {
	Send(text string, senderId, chatId uint64, sendingTime time.Time) error
	GetBatch(userId, chatId, pageNumber uint64) ([]core.Message, error)
}

type Repository struct {
	Contact
	Chat
	User
	Message
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		Contact: NewContactPostgres(db),
		Chat:    NewChatPostgres(db),
		User:    NewUserPostgres(db),
		Message: NewMessagePostgres(db),
	}
}
