package model

import (
	"fmt"
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

func (m *Messenger) CreateChat(id ID) {
	chat := &Chat{
		ID:        id,
		clients:   make(map[*Client]bool),
		broadcast: make(chan MessageSent),
	}
	m.chats[id] = chat
	chat.Run()
}

func (m *Messenger) Connect(conn *websocket.Conn, chatID ID, userID ID) error {
	client := newClient(conn, userID)

	chat, ok := m.chats[chatID]

	if !ok {
		return fmt.Errorf("chat with id %d not found", chatID)
	}

	chat.Connect(client)

	return nil
}
