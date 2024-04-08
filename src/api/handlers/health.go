package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(context *gin.Context) {
	context.JSON(http.StatusOK, "Working!")
	return
}
