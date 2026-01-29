package web

import (
	"encoding/json"
	"net/http"

	"github.com/davidabx-dev/gestao-estoque/internal/usecase"
)

type ProductHandler struct {
	CreateProductUseCase *usecase.CreateProductUseCase
	ListProductsUseCase  *usecase.ListProductsUseCase // <--- NOVO
}

// Atualizamos o construtor para receber os dois UseCases
func NewProductHandler(create *usecase.CreateProductUseCase, list *usecase.ListProductsUseCase) *ProductHandler {
	return &ProductHandler{
		CreateProductUseCase: create,
		ListProductsUseCase:  list,
	}
}

// Função CREATE (Mantivemos igual)
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateProductInputDto
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := h.CreateProductUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

// Função LIST (Nova!)
func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	output, err := h.ListProductsUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}