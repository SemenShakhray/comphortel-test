package router

import (
	"comphortel-test/internal/delivery/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handlers.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("users/:id", h.GetUser)

	return r
}
