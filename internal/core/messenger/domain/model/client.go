package model

import (
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type Client struct {
	conn           *websocket.Conn
	wg             *sync.WaitGroup
	send           chan string
	maxMessageSize int64
	pongWaitTime   time.Duration // Время ожидания Pong от клиента
	pingPeriodTime time.Duration // Интервал отправки Ping от сервера
}

func (c *Client) readPump(queue chan<- string) {
	defer c.wg.Done()

	c.conn.SetReadLimit(c.maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(c.pongWaitTime))
	c.conn.SetPongHandler(func(string) error { _ = c.conn.SetReadDeadline(time.Now().Add(c.pongWaitTime)); return nil })

	for {
		_, message, err := c.conn.ReadMessage()

		if err != nil {
			break
		}
		queue <- string(message)
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
			_ = c.conn.WriteMessage(websocket.TextMessage, []byte(message))

		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(c.pingPeriodTime))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// Run Запускает блокирующий ввод/вывод над клиентом
func (c *Client) Run(messageQueue chan<- string) {
	c.wg.Add(2)
	go c.readPump(messageQueue)
	go c.writePump()
	c.wg.Wait()
}

func newClient(conn *websocket.Conn) *Client {
	return &Client{
		conn:           conn,
		wg:             &sync.WaitGroup{},
		send:           make(chan string),
		pongWaitTime:   5 * time.Second,
		pingPeriodTime: 4 * time.Second,
		maxMessageSize: 10000,
	}
}
