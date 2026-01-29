package entity

import (
	"errors"
	"github.com/google/uuid" // Vamos precisar gerar IDs únicos
)

// Product define a estrutura do nosso dado
type Product struct {
	ID    string
	Name  string
	Price float64
}

// NewProduct é a nossa "fábrica". Ela valida as regras ANTES de criar.
func NewProduct(name string, price float64) (*Product, error) {
	// Validação 1: Nome não pode ser vazio
	if name == "" {
		return nil, errors.New("name is required")
	}
	
	// Validação 2: Preço deve ser maior que zero
	if price <= 0 {
		return nil, errors.New("price must be greater than zero")
	}

	// Se passou, cria o objeto na memória
	return &Product{
		ID:    uuid.New().String(), // Gera um ID aleatório único
		Name:  name,
		Price: price,
	}, nil
}