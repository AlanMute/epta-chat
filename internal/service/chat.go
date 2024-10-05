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
	return s.repo.Add(name, isDirect, ownerId, s.getUniqueMembers(members))
}

func (s *ChatService) Delete(userId, chatId uint64) error {
	return s.repo.Delete(userId, chatId)
}

func (s *ChatService) GetById(userId, chatId uint64) (core.Chat, error) {
	return s.repo.GetById(userId, chatId)
}

func (s *ChatService) GetAll(userId uint64) ([]core.Chat, error) {
	return s.repo.GetAll(userId)
}

func (s *ChatService) GetMembers(userId, chatId uint64) ([]core.UserInfo, error) {
	return s.repo.GetMembers(userId, chatId)
}

func (s *ChatService) getUniqueMembers(members []uint64) []uint64 {
	uniqMembMap := make(map[uint64]struct{}, len(members))
	for _, id := range members {
		uniqMembMap[id] = struct{}{}
	}

	if len(members) == len(uniqMembMap) {
		return members
	}

	members = members[:0]
	for id := range uniqMembMap {
		members = append(members, id)
	}

	return members
}
