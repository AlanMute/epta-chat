package service

import (
	"fmt"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
)

type ChatService struct {
	repo repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) Add(name string, isDirect bool, ownerId uint64, members []uint64) (uint64, error) {
	members = s.getUniqueMembers(append(members, ownerId))

	if isDirect && len(members) != 2 {
		return 0, fmt.Errorf("wrong amount of members for direct chat")
	}

	return s.repo.Add(name, isDirect, ownerId, members)
}

func (s *ChatService) AddMember(ownerId, chatId uint64, members []uint64) error {
	return s.repo.AddMember(ownerId, chatId, members)
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
