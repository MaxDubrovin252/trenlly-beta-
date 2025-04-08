package handler

import (
	"log/slog"
	"trenlly"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) SignUp(c *gin.Context) {
	var req trenlly.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return

	}

	id, err := h.services.Authorization.CreateUser(req)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logrus.Info("new user add", req.Username)
	c.JSON(200, gin.H{"id": id})

}

type SignInReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignIn(c *gin.Context) {
	var req SignInReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := h.services.Authorization.GenerateToken(req.Username, req.Password)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logrus.Info("user war authorize")
	logrus.Info(slog.String("username", req.Username))
	c.JSON(200, gin.H{
		"token": token,
	})
}
