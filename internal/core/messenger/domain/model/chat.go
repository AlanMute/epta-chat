package model

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"time"
)

type Chat struct {
	ID          ID
	clients     map[*Client]bool
	broadcast   chan MessageSent
	messageRepo repository.Message
}

func (c *Chat) startBroadcasting() {
	for message := range c.broadcast {
		message.ChatID = c.ID
		message.SendingTime = time.Now().Format(core.TimeFormat)

		_ = c.messageRepo.Send(message.Text, uint64(message.SenderID), uint64(c.ID), time.Now())

		for client := range c.clients {
			client.send <- message
		}
	}
}

func (c *Chat) Run() {
	go c.startBroadcasting()
}

func (c *Chat) Connect(client *Client) {
	// Добавление клиента в список подключений
	c.clients[client] = true

	// Блокировка на чтение и запись
	client.Run(c.broadcast)

	// Удаление клиента из списка подключений, так как соединение прервано
	delete(c.clients, client)
}
