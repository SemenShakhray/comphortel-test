package handlers

import (
	"comphortel-test/internal/models"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Servicer
}

func NewHandler(serv Servicer) Handler {
	return Handler{
		service: serv,
	}
}

type Servicer interface {
	GetUser(id string) (*models.User, error)
}

func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.service.GetUser(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, user)
}
