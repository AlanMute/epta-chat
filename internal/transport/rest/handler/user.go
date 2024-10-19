package handler

import (
	"net/http"
	"strconv"

	"github.com/KrizzMU/coolback-alkol/pkg/api/resp"

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
// @Router /user/sign-up [post]
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

	err = h.services.User.SignUp(info.Login, info.Password)
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
// @Router /user/sign-in [post]
// @Success 200 "Вход выполнен"
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) SignIn(c *gin.Context) {
	var (
		info   Sign
		userId uint64
		token  core.Tokens
		err    error
	)

	if err = c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.Error(err.Error()))
		return
	}

	userId, token, err = h.services.User.SignIn(info.Login, info.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	type resp struct {
		UserId uint64
		Token  core.Tokens
	}

	c.JSON(http.StatusOK, resp{
		UserId: userId,
		Token:  token,
	})
}

// SignIn godoc
// @Summary Войти
// @Description Войти
// @Tags User
// @Accept json
// @Produce json
// @Param body body Sign true "Данные для регистрации"
// @Router /user/sign-in [post]
// @Success 200 "Вход выполнен"
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) SetUsername(c *gin.Context) {
	var (
		info   UserName
		userId int
		err    error
	)

	userId, err = strconv.Atoi(c.Param("user-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("Invalid user id"))
		return
	}

	if err = c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.Error(err.Error()))
		return
	}

	err = h.services.User.SetUserName(uint64(userId), info.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

// Refresh godoc
// @Summary Обновить токены
// @Description Обновить токены
// @Tags User
// @Accept json
// @Produce json
// @Param body body Refresh true "Данные для регистрации"
// @Router /user/refresh [post]
// @Success 200 "Токены обновлены"
// @Failure 400 {object} resp.ErrorResponse "Запрос не правильно составлен"
// @Failure 500 {object} resp.ErrorResponse "Возникла внутренняя ошибка"
func (h *Handler) Refresh(c *gin.Context) {
	var (
		info  Refresh
		token string
		err   error
	)

	if err = c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.Error(err.Error()))
		return
	}

	token, err = h.services.User.Refresh(info.UserId, info.Token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, token)
}
