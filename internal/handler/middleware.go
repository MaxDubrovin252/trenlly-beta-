package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	UserCTX             = "UserId"
	AuthorizationHeader = "authorization"
)

func (h *Handler) UserIdentity(c *gin.Context) {

	header := c.GetHeader(AuthorizationHeader)

	if header == "" {
		c.JSON(400, gin.H{"auth": "empty token pole"})

		logrus.Info("failed auth:empty jwt token")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"auth": "invalid auth header"})
	}

	userId, err := h.services.ParseToken(headerParts[1])

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Set(UserCTX, userId)
}

func GetUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(UserCTX)

	if !ok {
		c.JSON(400, gin.H{"err": "user id not found"})
		return 0, nil

	}

	IdInt, ok := id.(int)

	if !ok {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return 0, nil
	}

	return IdInt, nil
}
