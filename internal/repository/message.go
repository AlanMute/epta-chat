package repository

import (
	"time"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

const batchSize = 100

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

func (r *MessageRepo) GetBatch(userId, chatId, pageNumber uint64) ([]core.Message, error) {
	var chat core.Chat

	if err := r.db.Where("id = ?", chatId).First(&chat).Error; err != nil {
		return nil, err
	}

	var member core.ChatMembers

	if err := r.db.Where("member_id = ? AND chat_id = ?", userId, chatId).First(&member).Error; err != nil {
		return nil, err
	}

	var messages []core.Message

	if err := r.db.Where("chat_id = ?", chatId).
		Limit(batchSize).
		Offset(batchSize * pageNumber).
		Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}
