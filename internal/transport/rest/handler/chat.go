package handler

import (
	"net/http"
	"strconv"

	"github.com/KrizzMU/coolback-alkol/pkg/api/resp"

	"github.com/gin-gonic/gin"
)

// GetChats godoc
// @Summary Получить список чатов пользователя
// @Description Получить список чатов пользователя
// @Security BearerAuth
// @Tags Chat
// @Accept json
// @Produce json
// @Router /chat/all [get]
// @Success 200 {object} []core.Chat
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) GetChats(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user-id")) //TODO: need will check token
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid user id"))
		return
	}

	chats, err := h.services.Chat.GetAll(uint64(userId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, chats)
}

// GetChatById godoc
// @Summary Получить чат по ID
// @Description Получить чат по ID
// @Security BearerAuth
// @Tags Chat
// @Param id path int true "ID чата"
// @Accept json
// @Produce json
// @Router /chat/{id} [get]
// @Success 200 {object} core.Chat
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) GetChatById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user-id")) //TODO: need will check token
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid user id"))
		return
	}

	chatId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid chat id"))
		return
	}

	chat, err := h.services.Chat.GetById(uint64(userId), uint64(chatId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, chat)
}

// GetChatMembers godoc
// @Summary Получить список участников чата
// @Description Получить список участников чата
// @Security BearerAuth
// @Tags Chat
// @Param id path int true "ID чата"
// @Accept json
// @Produce json
// @Router /chat/members/{id} [get]
// @Success 200 {object} []core.UserInfo
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) GetChatMembers(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user-id")) //TODO: need will check token
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid user id"))
		return
	}

	chatId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid chat id"))
		return
	}

	members, err := h.services.Chat.GetMembers(uint64(userId), uint64(chatId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, members)
}

// AddChat godoc
// @Summary Создать чат
// @Description Создать чат
// @Security BearerAuth
// @Tags Chat
// @Param body body AddChat true "Данные для создания чата"
// @Accept json
// @Produce json
// @Router /chat [post]
// @Success 201 "Чат создан"
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) AddChat(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid user id"))
		return
	}

	var info AddChat
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.Error(err.Error()))
		return
	}

	id, err := h.services.Chat.Add(info.Name, info.IsDirect, uint64(userId), info.Members)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	h.messengerService.CreateChat(id)

	c.Status(http.StatusCreated)
}

// AddMember godoc
// @Summary Добавить участника
// @Description Добавить участника
// @Security BearerAuth
// @Tags Chat
// @Param user-id query int true "ID пользователя"
// @Param body body AddMember true "Список users_id"
// @Accept json
// @Produce json
// @Router /chat/add/members [post]
// @Success 201 "Чат создан"
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) AddMember(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid user id"))
		return
	}

	var info AddMember
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.Error(err.Error()))
		return
	}

	err = h.services.Chat.AddMember(uint64(userId), info.ChatId, info.Members)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.Status(http.StatusCreated)
}

// DeleteChat godoc
// @Summary Удалить чат
// @Description Удалить чат
// @Security BearerAuth
// @Tags Chat
// @Param id path int true "ID чата"
// @Accept json
// @Produce json
// @Router /chat/{id} [delete]
// @Success 200 "Чат удален"
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) DeleteChat(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid user id"))
		return
	}

	chatId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid chat id"))
		return
	}

	err = h.services.Chat.Delete(uint64(userId), uint64(chatId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

// GetChatMessages godoc
// @Summary Получить историю сообщений
// @Description История сообщений получается постранично по 100 сообщений
// @Security BearerAuth
// @Tags Chat
// @Param chat-id query int true "ID чата"
// @Param page-id query int true "номер страницы"
// @Accept json
// @Produce json
// @Router /chat/{chat-id}/messages/{page-id} [get]
// @Success 200 {object} []core.Message
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) GetChatMessages(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid user id"))
		return
	}

	chatId, err := strconv.Atoi(c.Query("chat-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid chat id"))
		return
	}

	pageId, err := strconv.Atoi(c.Query("page-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid page id"))
		return
	}

	messages, err := h.services.Message.GetBatch(uint64(userId), uint64(chatId), uint64(pageId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, messages)
}
