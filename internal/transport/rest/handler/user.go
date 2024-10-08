package handler

import (
	"net/http"
	"strconv"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var (
		info Sign
		err  error
	)

	if err = c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.SignUp(info.Login, info.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) SignIn(c *gin.Context) {
	var (
		info  Sign
		token core.Tokens
		err   error
	)

	if err = c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err = h.services.SignIn(info.Login, info.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}

func (h *Handler) Refresh(c *gin.Context) {
	var (
		userId int
		info   Refresh
		token  string
		err    error
	)

	userId, err = strconv.Atoi(c.Param("user-id")) //TODO: need will check token
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = c.ShouldBindJSON(&info); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err = h.services.Refresh(uint64(userId), info.Token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}
