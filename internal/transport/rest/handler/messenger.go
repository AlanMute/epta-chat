package handler

import (
	"github.com/KrizzMU/coolback-alkol/pkg/api/resp"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Connect godoc
// @Summary Подключиться к мессенджеру
// @Description Установить websocket соединение с мессенджером
// @Tags Messenger
// @Param chat-id query int true "ID чата подключения"
// @Param user-id query int true "ID пользователя"
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

	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Chat ID should be int"))
		return
	}

	userIDStr := c.Query("user-id")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, resp.Error("User ID is required"))
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("User ID should be int"))
		return
	}

	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error("Unable to upgrade connection to websocket"))
		return
	}

	err = h.messenger.Connect(conn, chatID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error("Unable to connect to websocket"))
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
