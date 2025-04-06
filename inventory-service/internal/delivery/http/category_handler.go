package http

import (
	"net/http"
	"strconv"

	"inventory-service/internal/dto"
	"inventory-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	useCase usecase.CategoryUseCase
}

func NewCategoryHandler(router *gin.Engine, useCase usecase.CategoryUseCase) {
	handler := &CategoryHandler{useCase: useCase}

	categoryRoutes := router.Group("/categories")
	{
		categoryRoutes.POST("/", handler.CreateCategory)
		categoryRoutes.GET("/:id", handler.GetCategoryByID)
		categoryRoutes.PUT("/:id", handler.UpdateCategory)
		categoryRoutes.DELETE("/:id", handler.DeleteCategory)
		categoryRoutes.GET("/", handler.ListCategories)
	}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	category := req.ToDomain()
	if err := h.useCase.CreateCategory(c.Request.Context(), category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := h.useCase.GetCategoryByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if category == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	response := dto.NewCategoryResponse(*category)
	c.JSON(http.StatusOK, response)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	category := req.ToDomain()
	if err := h.useCase.UpdateCategory(c.Request.Context(), id, category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := h.useCase.DeleteCategory(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *CategoryHandler) ListCategories(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	// Parse filters from query parameters
	filter := map[string]interface{}{}
	for key, values := range c.Request.URL.Query() {
		if key != "limit" && key != "offset" {
			filter[key] = values[0]
		}
	}

	categories, err := h.useCase.ListCategories(c.Request.Context(), filter, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.CategoryResponse
	for _, category := range categories {
		response = append(response, dto.NewCategoryResponse(*category))
	}

	c.JSON(http.StatusOK, response)
}
