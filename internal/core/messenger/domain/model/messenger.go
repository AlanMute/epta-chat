package model

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
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
	go m.startBroadcasting()
	return m
}

func (m *Messenger) Connect(w http.ResponseWriter, r *http.Request) error {
	wg := sync.WaitGroup{}

	conn, err := m.upgrader.Upgrade(w, r, nil)

	if err != nil {
		return err
	}

	client := newClient(conn)
	m.clients[client] = true

	wg.Add(2)
	go client.StartReading(&wg, m.broadcast)
	go client.StartSending(&wg)

	wg.Wait()

	delete(m.clients, client)
	return nil
}
