package http

import (
	"encoding/json"
	"inventory-service/internal/dto"
	"inventory-service/internal/usecase"
	"net/http"
)

type ProductHandler struct {
	productUC usecase.ProductUseCase
}

func NewProductHandler(uc usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUC: uc}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	product := req.ToDomain()
	if err := h.productUC.CreateProduct(r.Context(), product); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	response := dto.NewProductResponse(product)
	writeJSON(w, http.StatusCreated, response)
}
