package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husfuu/go-order/helper"
	"github.com/husfuu/go-order/service"
)

type orderHandler struct {
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) *orderHandler {
	return &orderHandler{service}
}

func (h *orderHandler) GetOrders(c *gin.Context) {
	orders, err := h.service.GetOrders()

	if err != nil {
		response := helper.APIResponse("Error to get all orders", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of orders", http.StatusOK, "success", orders)

	c.JSON(http.StatusOK, response)
}
