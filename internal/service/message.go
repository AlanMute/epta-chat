package service

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
)

type MessageService struct {
	repo repository.Message
}

func NewMessageService(repo repository.Message) *MessageService {
	return &MessageService{
		repo: repo,
	}
}

func (s *MessageService) GetBatch(userId, chatId, pageNumber uint64) ([]core.MessageInfo, error) {
	messages, err := s.repo.GetBatch(userId, chatId, pageNumber)
	if err != nil {
		return nil, err
	}

	messagesInfo := make([]core.MessageInfo, len(messages))

	for i, message := range messages {
		messagesInfo[len(messages)-i-1] = core.MessageInfo{
			ID:          message.ID,
			Text:        message.Text,
			SendingTime: message.SendingTime,
			SenderId:    message.SenderId,
			UserName:    message.Sender.UserName,
			ChatId:      message.ChatId,
		}
	}

	return messagesInfo, nil
}
