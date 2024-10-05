package service

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
)

type ChatService struct {
	repo repository.Chat
}

func NewChatSevice(repo repository.Chat) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) Add(name string, isDirect bool, ownerId uint64, members []uint64) (uint64, error) {
	return 0, nil
}

func (s *ChatService) Delete(userId, chatId uint64) error {
	return nil
}

func (s *ChatService) GetById(userId, chatId uint64) (core.Chat, error) {
	return core.Chat{}, nil
}

func (s *ChatService) GetAll(userId uint64) ([]core.Chat, error) {
	return nil, nil
}

func (s *ChatService) GetMembers(userId, chatId uint64) ([]core.UserInfo, error) {
	return nil, nil
}
