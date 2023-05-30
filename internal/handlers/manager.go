package handlers

import (
	"github.com/gin-gonic/gin"
	"serviceB/internal/services/serviceB"
)

// Manager реализация api сервисов
type manager struct {
	ServiceB serviceB.IServiceB
}

type Manager interface {
	TransferHandler(c *gin.Context)
}

func NewManager(serviceB serviceB.IServiceB) Manager {
	return &manager{
		ServiceB: serviceB,
	}
}
