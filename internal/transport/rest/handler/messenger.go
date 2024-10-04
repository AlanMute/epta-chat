package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Connect godoc
// @Summary Подключиться к мессенджеру
// @Description Установить websocket соединение с мессенджером
// @Tags Messenger
// @Accept json
// @Produce json
// @Router /messenger/connect [get]
func (h *Handler) Connect(c *gin.Context) {
	err := h.messenger.Connect(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Connection Established"})
}
