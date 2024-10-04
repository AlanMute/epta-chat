package model

type Chat struct {
	ID        int
	clients   map[*Client]bool
	broadcast chan string
}

func (c *Chat) startBroadcasting() {
	for message := range c.broadcast {
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
