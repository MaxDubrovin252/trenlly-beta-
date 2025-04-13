package handler

import (
	"log/slog"
	"strconv"
	"trenlly"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateTren(c *gin.Context) {
	userId, err := GetUserID(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var req trenlly.Trening

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, result, err := h.services.Trening.CreateTren(req, userId)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"id":     id,
		"result": result,
	})

}
func (h *Handler) GetAllTren(c *gin.Context) {
	userId, err := GetUserID(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	trens, err := h.services.Trening.GetAll(userId)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if trens == nil {
		c.JSON(404, gin.H{"tren is empty": err})
	}
	logrus.Info("user was get all trens")
	logrus.Info(slog.Int("UserID", userId))

	c.JSON(200, trens)

}
func (h *Handler) GetTrenById(c *gin.Context) {
	userId, err := GetUserID(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	trenId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return

	}

	trenli, err := h.services.Trening.GetById(userId, trenId)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	logrus.Info("user get tren")
	logrus.Info(slog.Int("userId", userId))
	logrus.Info(slog.Int("treId", trenId))
	c.JSON(200, gin.H{
		"tren":   trenli,
		"STATUS": "OK",
	})

}

func (h *Handler) UpdateTren(c *gin.Context) {

	userId, err := GetUserID(c)

	if err != nil {
		c.JSON(40, gin.H{"error": err.Error()})
		return
	}

	trenId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var req trenlly.UpdateTrenItem

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Trening.UpdateTren(userId, trenId, req); err != nil {
		c.JSON(400, gin.H{"error": err})
		logrus.Errorf("error %s", err)
		return
	}

	logrus.Info("user was updated tren")
	logrus.Info(slog.Int("userId", userId))
	logrus.Info("trenId", trenId)

	c.JSON(200, gin.H{
		"STATUS": "updated",
	})

}
func (h *Handler) DeleteTren(c *gin.Context) {
	userId, err := GetUserID(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	trenId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Trening.Delete(userId, trenId); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"STATUS": "DELETED",
	})
}
