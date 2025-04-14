package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	categoryproto "api-gateway/proto/category"
	orderpb "api-gateway/proto/order"
	productproto "api-gateway/proto/product"
)

func main() {
	r := gin.Default()
	gin.SetMode(getEnv("GIN_MODE", "release"))

	// Connect to Order Service
	orderConn, err := grpc.Dial(getEnv("ORDER_SERVICE_GRPC", "localhost:8020"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to order service: %v", err)
	}
	defer orderConn.Close()
	orderClient := orderpb.NewOrderServiceClient(orderConn)

	// Connect to Inventory Service
	inventoryConn, err := grpc.Dial(getEnv("INVENTORY_SERVICE_GRPC", "localhost:8010"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to inventory service: %v", err)
	}
	defer inventoryConn.Close()
	ProductClient := productproto.NewProductServiceClient(inventoryConn)
	CategotyClient := categoryproto.NewCategoryServiceClient(inventoryConn)

	// Order Endpoints
	r.POST("/api/v1/orders", func(c *gin.Context) {
		var req orderpb.CreateOrderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := orderClient.CreateOrder(context.Background(), &req)
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/orders/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		res, err := orderClient.GetOrderByID(context.Background(), &orderpb.GetOrderRequest{Id: int32(id)})
		handleResponse(c, res, err)
	})

	r.POST("/api/v1/orders/:id/status", func(c *gin.Context) {
		var req orderpb.UpdateOrderStatusRequest
		id, _ := strconv.Atoi(c.Param("id"))
		req.Id = int32(id)
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := orderClient.UpdateOrderStatus(context.Background(), &req)
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/orders", func(c *gin.Context) {
		userID, _ := strconv.Atoi(c.Query("user_id"))
		limit := queryInt(c, "limit", 10)
		offset := queryInt(c, "offset", 0)
		res, err := orderClient.ListUserOrders(context.Background(), &orderpb.ListOrdersRequest{
			UserId: int32(userID), Limit: int32(limit), Offset: int32(offset),
		})
		handleResponse(c, res, err)
	})

	// Inventory Endpoints
	r.POST("/api/v1/products", func(c *gin.Context) {
		var req productproto.CreateProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := ProductClient.CreateProduct(context.Background(), &req)
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		res, err := ProductClient.GetProductByID(context.Background(), &productproto.GetProductRequest{Id: int32(id)})
		handleResponse(c, res, err)
	})

	r.DELETE("/api/v1/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := ProductClient.DeleteProduct(context.Background(), &productproto.DeleteProductRequest{Id: int32(id)})
		handleResponse(c, gin.H{"message": "deleted"}, err)
	})

	r.GET("/api/v1/products", func(c *gin.Context) {
		limit := queryInt(c, "limit", 10)
		offset := queryInt(c, "offset", 0)
		res, err := ProductClient.ListProducts(context.Background(), &productproto.ListProductsRequest{
			Limit: int32(limit), Offset: int32(offset),
		})
		handleResponse(c, res, err)
	})

	r.POST("/api/v1/categories", func(c *gin.Context) {
		var req categoryproto.CreateCategoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := CategotyClient.CreateCategory(context.Background(), &req)
		handleResponse(c, res, err)
	})

	r.GET("/api/v1/categories/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		res, err := CategotyClient.GetCategoryByID(context.Background(), &categoryproto.GetCategoryRequest{Id: int32(id)})
		handleResponse(c, res, err)
	})

	r.DELETE("/api/v1/categories/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := CategotyClient.DeleteCategory(context.Background(), &categoryproto.DeleteCategoryRequest{Id: int32(id)})
		handleResponse(c, gin.H{"message": "deleted"}, err)
	})

	r.GET("/api/v1/categories", func(c *gin.Context) {
		limit := queryInt(c, "limit", 10)
		offset := queryInt(c, "offset", 0)
		res, err := CategotyClient.ListCategories(context.Background(), &categoryproto.ListCategoriesRequest{
			Limit: int32(limit), Offset: int32(offset),
		})
		handleResponse(c, res, err)
	})

	// Start the server
	server := &http.Server{
		Addr:    ":" + getEnv("HTTP_PORT", "8000"),
		Handler: r,
	}

	go func() {
		log.Println("API Gateway running on", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	waitForShutdown(server)
}

func handleResponse(c *gin.Context, res any, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func waitForShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown failed: %v", err)
	}
	log.Println("Server exited gracefully")
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func queryInt(c *gin.Context, key string, defaultVal int) int {
	val, err := strconv.Atoi(c.Query(key))
	if err != nil {
		return defaultVal
	}
	return val
}
