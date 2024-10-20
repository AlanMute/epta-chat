package model

import (
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type Client struct {
	userID         ID
	userName       string
	conn           *websocket.Conn
	wg             *sync.WaitGroup
	send           chan MessageSent
	maxMessageSize int64
	pongWaitTime   time.Duration // Время ожидания Pong от клиента
	pingPeriodTime time.Duration // Интервал отправки Ping от сервера
}

func (c *Client) readPump(queue chan<- MessageSent) {
	defer c.wg.Done()

	c.conn.SetReadLimit(c.maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(c.pongWaitTime))
	c.conn.SetPongHandler(func(string) error { _ = c.conn.SetReadDeadline(time.Now().Add(c.pongWaitTime)); return nil })

	for {
		_, data, err := c.conn.ReadMessage()

		if err != nil {
			break
		}

		var message MessageReceived
		err = json.Unmarshal(data, &message)
		if err != nil {
			break
		}

		queue <- MessageSent{
			SenderID: c.userID,
			UserName: c.userName,
			Text:     message.Text,
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(c.pingPeriodTime)

	defer func() {
		ticker.Stop()
		_ = c.conn.Close()
		c.wg.Done()
	}()

	for {
		select {
		case message, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(c.pingPeriodTime))
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			data, err := json.Marshal(message)

			if err != nil {
				return
			}

			_ = c.conn.WriteMessage(websocket.TextMessage, data)

		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(c.pingPeriodTime))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// Run Запускает блокирующий ввод/вывод над клиентом
func (c *Client) Run(messageQueue chan<- MessageSent) {
	defer func() {
		_ = c.conn.Close()
	}()

	c.wg.Add(2)
	go c.readPump(messageQueue)
	go c.writePump()
	c.wg.Wait()
}

func (c *Client) Stop() {
	defer func() {
		_ = c.conn.Close()
	}()
}

func newClient(conn *websocket.Conn, userID ID, userName string) *Client {
	return &Client{
		userID:         userID,
		userName:       userName,
		conn:           conn,
		wg:             &sync.WaitGroup{},
		send:           make(chan MessageSent),
		pongWaitTime:   5 * time.Second,
		pingPeriodTime: 4 * time.Second,
		maxMessageSize: 10000,
	}
}
