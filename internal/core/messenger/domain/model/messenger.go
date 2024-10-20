package model

import (
	"fmt"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"github.com/gorilla/websocket"
)

type ID uint64

type Messenger struct {
	chats map[ID]*Chat
}

func NewMessenger() *Messenger {
	m := &Messenger{
		chats: make(map[ID]*Chat),
	}
	return m
}

func (m *Messenger) CreateChat(id ID, messageRepo repository.Message) {
	chat := &Chat{
		ID:          id,
		clients:     make(map[*Client]bool),
		broadcast:   make(chan MessageSent),
		messageRepo: messageRepo,
	}
	m.chats[id] = chat
	chat.Run()
}

func (m *Messenger) Connect(conn *websocket.Conn, chatID ID, userID ID, userName string) error {
	client := newClient(conn, userID, userName)

	chat, ok := m.chats[chatID]

	if !ok {
		client.Stop()
		return fmt.Errorf("chat with id %d not found", chatID)
	}

	chat.Connect(client)

	return nil
}
