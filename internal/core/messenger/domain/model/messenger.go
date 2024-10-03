package model

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type Messenger struct {
	upgrader  websocket.Upgrader
	broadcast chan string
	clients   map[*Client]bool
}

func (m *Messenger) startBroadcasting() {
	for message := range m.broadcast {
		for client := range m.clients {
			client.send <- message
		}
	}
}

func NewMessenger() *Messenger {
	m := &Messenger{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  4096,
			WriteBufferSize: 4096,
		},
		broadcast: make(chan string),
		clients:   make(map[*Client]bool),
	}
	return m
}

func (m *Messenger) Run() {
	go m.startBroadcasting()
}

func (m *Messenger) Connect(w http.ResponseWriter, r *http.Request) error {
	conn, err := m.upgrader.Upgrade(w, r, nil)

	if err != nil {
		return err
	}

	client := newClient(conn)

	// Добавление клиента в список подключений
	m.clients[client] = true

	// Блокировка на чтение и запись
	client.Run(m.broadcast)

	// Удаление клиента из списка подключений, так как соединение прервано
	delete(m.clients, client)

	return nil
}
