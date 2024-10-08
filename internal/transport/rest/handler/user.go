package handler

import (
	"github.com/KrizzMU/coolback-alkol/pkg/api/resp"
	"net/http"
	"strconv"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/gin-gonic/gin"
)

// SignUp godoc
// @Summary Зарегистрироваться
// @Description Зарегистрироваться
// @Tags User
// @Accept json
// @Produce json
// @Param body body Sign true "Данные для регистрации"
// @Router /sign-up [post]
// @Success 200 "Регистрация выполнена"
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) SignUp(c *gin.Context) {
	var (
		info Sign
		err  error
	)

	if err = c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.Error(err.Error()))
		return
	}

	err = h.services.SignUp(info.Login, info.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

// SignIn godoc
// @Summary Войти
// @Description Войти
// @Tags User
// @Accept json
// @Produce json
// @Param body body Sign true "Данные для регистрации"
// @Router /sign-in [post]
// @Success 200 "Вход выполнен"
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) SignIn(c *gin.Context) {
	var (
		info  Sign
		token core.Tokens
		err   error
	)

	if err = c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.Error(err.Error()))
		return
	}

	token, err = h.services.SignIn(info.Login, info.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, token)
}

// Refresh godoc
// @Summary Обновить токены
// @Description Обновить токены
// @Tags User
// @Accept json
// @Produce json
// @Param user-id query int true "ID пользователя"
// @Param body body Refresh true "Данные для регистрации"
// @Router /refresh [post]
// @Success 200 "Токены обновлены"
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) Refresh(c *gin.Context) {
	var (
		userId int
		info   Refresh
		token  string
		err    error
	)

	userId, err = strconv.Atoi(c.Param("user-id")) //TODO: need will check token
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid user id"))
		return
	}

	if err = c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.Error(err.Error()))
		return
	}

	token, err = h.services.Refresh(uint64(userId), info.Token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, token)
}
