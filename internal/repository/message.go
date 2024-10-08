package repository

import (
	"time"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type MessageRepo struct {
	db *gorm.DB
}

func NewMessagePostgres(db *gorm.DB) *MessageRepo {
	return &MessageRepo{
		db: db,
	}
}

func (r *MessageRepo) Send(text string, senderId, chatId uint64, sendingTime time.Time) error {
	message := core.Message{
		Text:        text,
		SendingTime: sendingTime,
		SenderId:    senderId,
		ChatId:      chatId,
	}

	if result := r.db.Create(&message); result.Error != nil {
		return result.Error
	}

	return nil
}
