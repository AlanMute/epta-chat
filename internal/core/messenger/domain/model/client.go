package model

import (
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type Client struct {
	conn           *websocket.Conn
	send           chan string
	maxMessageSize int64
	pongWaitTime   time.Duration
	pingPeriodTime time.Duration
}

func (c *Client) StartReading(wg *sync.WaitGroup, queue chan<- string) {
	defer wg.Done()
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

func (c *Client) StartSending(wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(c.pingPeriodTime)
	defer func() {
		ticker.Stop()
		_ = c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(c.pingPeriodTime))
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)

			if err != nil {
				return
			}

			_, _ = w.Write([]byte(message))

		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(c.pingPeriodTime))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func newClient(conn *websocket.Conn) *Client {
	return &Client{
		conn:           conn,
		send:           make(chan string),
		pongWaitTime:   5 * time.Second,
		pingPeriodTime: 4 * time.Second,
		maxMessageSize: 10000,
	}
}
