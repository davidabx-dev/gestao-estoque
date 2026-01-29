package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	// Cenário 1: Tudo certo
	p, err := NewProduct("Notebook Gamer", 5000.0)

	assert.Nil(t, err)           // Não deve ter erro
	assert.NotNil(t, p)          // O produto deve existir
	assert.Equal(t, "Notebook Gamer", p.Name)
	assert.Equal(t, 5000.0, p.Price)
}

func TestNewProductWhenPriceIsInvalid(t *testing.T) {
	// Cenário 2: Preço zero (Erro proposital)
	p, err := NewProduct("Notebook Gamer", 0)

	assert.NotNil(t, err)        // DEVE retornar um erro
	assert.EqualError(t, err, "price must be greater than zero")
	assert.Nil(t, p)             // O produto não deve ser criado
}