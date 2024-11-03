package handler

import (
	"net/http"
	"strconv"

	"github.com/KrizzMU/coolback-alkol/pkg/api/resp"
	"github.com/gin-gonic/gin"
)

// Connect godoc
// @Summary Подключиться к мессенджеру
// @Description Установить websocket соединение с чатом. Чтобы отправить сообщение в чат нужно сформировать json в формате { text: string }, приходить сообщения буду в формате { id: uint, text: string, senderId: uint, userName: string, chatId: uint, sendingTime: string }
// @Security BearerAuth
// @Tags Messenger
// @Param chat-id query int true "ID чата подключения"
// @Accept json
// @Produce json
// @Router /messenger/connect [get]
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла непредвиденная ошибка"
func (h *Handler) Connect(c *gin.Context) {
	chatIDStr := c.Query("chat-id")
	if chatIDStr == "" {
		c.JSON(http.StatusBadRequest, resp.Error("Chat ID is required"))
		return
	}

	chatID, err := strconv.ParseUint(chatIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Chat ID should be uint64"))
		return
	}

	userIDStr := c.Param("user-id")
	if userIDStr == "" {
		c.JSON(http.StatusUnauthorized, resp.Error("User ID is required"))
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("User ID should be uint64"))
		return
	}

	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error("Unable to upgrade connection to websocket"))
		return
	}

	err = h.messengerService.JoinChat(conn, userID, chatID)
	if err != nil {
		_ = conn.Close()
		c.JSON(http.StatusInternalServerError, resp.Error("Unable to connect to websocket"))
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
