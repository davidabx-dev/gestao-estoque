package usecase

import (
	"github.com/davidabx-dev/gestao-estoque/internal/entity"
)

// 1. InputDTO: O que eu preciso receber para trabalhar?
type CreateProductInputDto struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// 2. OutputDTO: O que eu vou devolver para quem chamou?
type CreateProductOutputDto struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// 3. A Estrutura do Use Case
type CreateProductUseCase struct {
	// Ele precisa de um Repositório (Interface) para salvar depois
	ProductRepository entity.ProductRepository
}

// 4. Construtor: Cria o Use Case já injetando a dependência (Repository)
func NewCreateProductUseCase(productRepository entity.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		ProductRepository: productRepository,
	}
}

// 5. Execute: Onde a mágica acontece
func (u *CreateProductUseCase) Execute(input CreateProductInputDto) (*CreateProductOutputDto, error) {
	// A. Cria a Entidade (Valida as regras: preço > 0, nome required)
	product, err := entity.NewProduct(input.Name, input.Price)
	if err != nil {
		return nil, err
	}

	// B. Salva no Banco (Usando a interface, sem saber qual banco é)
	err = u.ProductRepository.Create(product)
	if err != nil {
		return nil, err
	}

	// C. Retorna o OutputDTO
	return &CreateProductOutputDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}