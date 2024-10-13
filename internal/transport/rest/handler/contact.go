package handler

import (
	"net/http"
	"strconv"

	"github.com/KrizzMU/coolback-alkol/pkg/api/resp"

	"github.com/gin-gonic/gin"
)

// GetContacts godoc
// @Summary Получить список контактов пользователя
// @Description Получить список контактов пользователя
// @Security BearerAuth
// @Tags Contact
// @Accept json
// @Produce json
// @Router /contact/all [get]
// @Success 200 {object} []core.UserInfo
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) GetContacts(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user-id")) //TODO: need will check token
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid user id"))
		return
	}

	contacts, err := h.services.Contact.GetAll(uint64(userId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, contacts)
}

// GetContactById godoc
// @Summary Получить контакт по ID
// @Description Получить контакт по ID
// @Security BearerAuth
// @Tags Contact
// @Param id path int true "ID контакта"
// @Accept json
// @Produce json
// @Router /contact/{id} [get]
// @Success 200 {object} core.UserInfo
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) GetContactById(c *gin.Context) {
	contactId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid contact id"))
		return
	}

	contact, err := h.services.Contact.GetById(uint64(contactId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, contact)
}

// AddContact godoc
// @Summary Создать контакт
// @Description Создать контакт
// @Security BearerAuth
// @Tags Contact
// @Accept json
// @Produce json
// @Router /contact [post]
// @Success 201 "Контакт создан"
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) AddContact(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user-id")) //TODO: need will check token
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid user id"))
		return
	}

	var info AddContact
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.Error(err.Error()))
		return
	}

	err = h.services.Contact.Add(uint64(userId), info.ContactId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.Status(http.StatusCreated)
}

// DeleteContact godoc
// @Summary Удалить контакт
// @Description Удалить контакт
// @Security BearerAuth
// @Tags Contact
// @Param id path int true "ID контакта"
// @Accept json
// @Produce json
// @Router /contact [delete]
// @Success 200 "Контакт удален"
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) DeleteContact(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user-id")) //TODO: need will check token
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid user id"))
		return
	}

	contactId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid contact id"))
		return
	}

	err = h.services.Contact.Delete(uint64(userId), uint64(contactId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}
