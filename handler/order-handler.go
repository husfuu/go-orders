package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husfuu/go-order/dao"
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

func (h *orderHandler) GetOrder(c *gin.Context) {
	var input dao.GetOrderDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("Failed to get detail of order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	orderDetail, err := h.service.GetOrderById(input)

	if err != nil {
		response := helper.APIResponse("Failed to get detail of order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign detail", http.StatusOK, "success", orderDetail)
	c.JSON(http.StatusOK, response)
}

func (h *orderHandler) CreateOrder(c *gin.Context) {
	var input dao.CreateOrderInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newOrder, err := h.service.CreateOrder(input)
	if err != nil {
		response := helper.APIResponse("Failed to create order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create order", http.StatusOK, "success", newOrder)
	c.JSON(http.StatusOK, response)
}
