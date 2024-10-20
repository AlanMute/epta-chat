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

func (s *MessageService) GetBatch(userId, chatId, pageNumber uint64) ([]core.Message, error) {
	return s.repo.GetBatch(userId, chatId, pageNumber)
}
