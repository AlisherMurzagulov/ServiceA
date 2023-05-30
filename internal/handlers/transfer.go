package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"serviceB/internal/models"
)

// TransferHandler
// @Summary Метод для добавление стакана по фин инструменту
// @Tags OrderBooksCash
// @Accept  json
// @Produce  json
// @Param comparisons body models.Request{} false "Тело запроса для сдобавление стакан по фин инструменту"
// @Success 200 {object} models.BaseResponse{} "Статус добавление стакан по фин инструменту"
// @Router /api/v1/order-book/add [post]
func (m *manager) TransferHandler(c *gin.Context) {
	var request models.Request
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ShouldBindJSON": "Invalid request data"})
		return
	}
	err = m.ServiceB.Transfer(c, request, "Alisher")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Transfer": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transfer request sent"})
}
