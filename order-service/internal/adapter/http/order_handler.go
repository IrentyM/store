package handler

// import (
// 	"net/http"
// 	"strconv"

// 	"order-service/internal/dto"

// 	"github.com/gin-gonic/gin"
// )

// type OrderHandler interface {
// 	CreateOrder(c *gin.Context)
// 	GetOrderByID(c *gin.Context)
// 	UpdateOrder(c *gin.Context)
// 	DeleteOrder(c *gin.Context)
// 	ListOrders(c *gin.Context)
// }

// type orderHandler struct {
// 	useCase OrderUseCase
// }

// func NewOrderHandler(useCase OrderUseCase) OrderHandler {
// 	return &orderHandler{useCase: useCase}
// }

// func (h *orderHandler) CreateOrder(c *gin.Context) {
// 	var req dto.CreateOrderRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
// 		return
// 	}

// 	order, items := req.ToDomain()
// 	orderID, err := h.useCase.CreateOrder(c.Request.Context(), order, items)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"order_id": orderID})
// }

// func (h *orderHandler) GetOrderByID(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
// 		return
// 	}

// 	order, items, err := h.useCase.GetOrderByID(c.Request.Context(), id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if order == nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
// 		return
// 	}

// 	response := dto.NewOrderResponse(*order, items)
// 	c.JSON(http.StatusOK, response)
// }

// func (h *orderHandler) UpdateOrder(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
// 		return
// 	}

// 	var req dto.CreateOrderRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
// 		return
// 	}

// 	order, _ := req.ToDomain()
// 	if err := h.useCase.UpdateOrder(c.Request.Context(), id, order); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.Status(http.StatusOK)
// }

// func (h *orderHandler) DeleteOrder(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
// 		return
// 	}

// 	if err := h.useCase.DeleteOrder(c.Request.Context(), id); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.Status(http.StatusNoContent)
// }

// func (h *orderHandler) ListOrders(c *gin.Context) {
// 	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
// 	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

// 	// Parse filters from query parameters
// 	filter := map[string]interface{}{}
// 	for key, values := range c.Request.URL.Query() {
// 		if key != "limit" && key != "offset" {
// 			filter[key] = values[0]
// 		}
// 	}

// 	orders, err := h.useCase.ListOrders(c.Request.Context(), filter, limit, offset)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, orders)
// }
