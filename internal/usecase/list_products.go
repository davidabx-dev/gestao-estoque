package usecase

import "github.com/davidabx-dev/gestao-estoque/internal/entity"

// 1. OutputDTO: O que vamos devolver para o usuário?
type ListProductsOutputDto struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// 2. Struct do UseCase
type ListProductsUseCase struct {
	ProductRepository entity.ProductRepository
}

// 3. Construtor
func NewListProductsUseCase(productRepository entity.ProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{
		ProductRepository: productRepository,
	}
}

// 4. Execute
func (u *ListProductsUseCase) Execute() ([]*ListProductsOutputDto, error) {
	// Chama o repositório (Banco)
	products, err := u.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	// Converte as Entidades para DTOs (Boas práticas: não devolva a Entidade pura)
	var productsDto []*ListProductsOutputDto
	for _, p := range products {
		productsDto = append(productsDto, &ListProductsOutputDto{
			ID:    p.ID,
			Name:  p.Name,
			Price: p.Price,
		})
	}

	return productsDto, nil
}