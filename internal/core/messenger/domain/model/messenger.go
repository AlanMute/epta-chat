package model

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

type Messenger struct {
	upgrader websocket.Upgrader
	chats    map[int]*Chat
}

func NewMessenger() *Messenger {
	m := &Messenger{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  4096,
			WriteBufferSize: 4096,
		},
		chats: make(map[int]*Chat),
	}
	return m
}

func (m *Messenger) CreateChat(id int) {
	chat := &Chat{
		ID:        id,
		clients:   make(map[*Client]bool),
		broadcast: make(chan string),
	}
	m.chats[id] = chat
	chat.Run()
}

func (m *Messenger) Connect(w http.ResponseWriter, r *http.Request) error {
	chatIDStr := r.URL.Query().Get("chat_id")

	if chatIDStr == "" {
		return fmt.Errorf("chat_id is empty")
	}

	chatID, err := strconv.Atoi(chatIDStr)

	if err != nil {
		return err
	}

	conn, err := m.upgrader.Upgrade(w, r, nil)

	if err != nil {
		return err
	}

	client := newClient(conn)

	chat, ok := m.chats[chatID]

	if !ok {
		return fmt.Errorf("chat with id %d not found", chatID)
	}

	chat.Connect(client)

	return nil
}
