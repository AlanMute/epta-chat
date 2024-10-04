package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Connect godoc
// @Summary Подключиться к мессенджеру
// @Description Установить websocket соединение с мессенджером
// @Tags Messenger
// @Accept json
// @Produce json
// @Router /messenger/connect [get]
func (h *Handler) Connect(c *gin.Context) {
	chatIDStr := c.Query("chat_id")
	if chatIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No chat ID in request context"})
		return
	}

	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chat ID should be int"})
		return
	}

	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to upgrade connection to websocket"})
		return
	}

	err = h.messenger.Connect(conn, chatID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to connect to websocket"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
