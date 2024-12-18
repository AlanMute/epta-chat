package handler

import (
	"net/http"
	"strconv"

	"github.com/KrizzMU/coolback-alkol/pkg/api/resp"
	"github.com/jinzhu/gorm"

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
	userId, err := strconv.Atoi(c.Param(userIdParam))
	if err != nil {
		c.JSON(http.StatusUnauthorized, resp.Error(invalidUserId))
		return
	}

	contacts, err := h.services.Contact.GetAll(uint64(userId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	if len(contacts) == 0 {
		c.JSON(http.StatusOK, [0]uint32{})
	} else {
		c.JSON(http.StatusOK, contacts)
	}
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
// @Param body body AddContact true "Логин контакта"
// @Router /contact [post]
// @Success 201 "Контакт создан"
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) AddContact(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param(userIdParam))
	if err != nil {
		c.JSON(http.StatusUnauthorized, resp.Error(invalidUserId))
		return
	}

	var info AddContact
	if err := c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.Error(err.Error()))
		return
	}

	err = h.services.Contact.Add(uint64(userId), info.ContactLogin)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithStatusJSON(http.StatusBadRequest, resp.Error("user not found"))
			return
		}
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
// @Router /contact/{id} [delete]
// @Success 200 "Контакт удален"
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) DeleteContact(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param(userIdParam))
	if err != nil {
		c.JSON(http.StatusUnauthorized, resp.Error(invalidUserId))
		return
	}

	contactId, err := strconv.Atoi(c.Param(idParam))
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
