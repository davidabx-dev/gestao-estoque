package entity

// ProductRepository é o CONTRATO.
// Quem quiser ser o banco de dados desse sistema (Seja Postgres, MySQL ou Fake)
// TEM QUE ter esses métodos.
type ProductRepository interface {
	Create(product *Product) error
	FindAll() ([]*Product, error)
}