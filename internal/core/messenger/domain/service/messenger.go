package service

import (
	"errors"
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/core/messenger/domain/model"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"github.com/gorilla/websocket"
	"slices"
)

type Messenger struct {
	chatRepo    repository.Chat
	userRepo    repository.User
	messageRepo repository.Message
	messenger   *model.Messenger
}

func NewMessenger(
	chatRepo repository.Chat,
	userRepo repository.User,
	messageRepo repository.Message,
	messenger *model.Messenger,
) (*Messenger, error) {
	m := &Messenger{
		chatRepo:    chatRepo,
		userRepo:    userRepo,
		messenger:   messenger,
		messageRepo: messageRepo,
	}
	chatsIds, err := chatRepo.FetchAllChatIDs()
	if err != nil {
		return m, err
	}

	for _, chatId := range chatsIds {
		m.CreateChat(chatId)
	}

	return m, nil
}

func (s *Messenger) JoinChat(conn *websocket.Conn, userID, chatID uint64) error {
	members, err := s.chatRepo.GetMembers(userID, chatID)

	if err != nil {
		return err
	}

	index := slices.IndexFunc(members, func(user core.UserInfo) bool {
		return user.ID == userID
	})

	if index == -1 {
		return errors.New("user is not member of that chat")
	}

	user, err := s.userRepo.GetById(userID)

	if err != nil {
		return err
	}

	err = s.messenger.Connect(conn, model.ID(chatID), model.ID(userID), user.UserName)

	return err
}

func (s *Messenger) CreateChat(chatID uint64) {
	s.messenger.CreateChat(model.ID(chatID), s.messageRepo)
}
