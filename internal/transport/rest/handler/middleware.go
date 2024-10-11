package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) isLogedIn(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Empty auth header!")
		return
	}

	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid auth header!")
		return
	}

	if len(parts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Tocken is empty!")
		return
	}

	id, err := h.tokenManger.Parse(parts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if id != "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Access denied")
		return
	}
}
