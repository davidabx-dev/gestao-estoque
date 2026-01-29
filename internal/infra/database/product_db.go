package database

import (
	"database/sql"
	"github.com/davidabx-dev/gestao-estoque/internal/entity"
)

// ProductDB é a struct que vai conversar com o banco.
// Ela guarda a conexão (*sql.DB) dentro dela.
type ProductDB struct {
	DB *sql.DB
}

// NewProductDB cria uma nova instância do nosso banco
func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{DB: db}
}

// Create insere o produto na tabela.
// Esse método CASA perfeitamente com a interface ProductRepository que criamos antes.
func (p *ProductDB) Create(product *entity.Product) error {
	// Query SQL crua (Mostra que você sabe SQL, que a vaga pede)
	stmt, err := p.DB.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Executa o comando
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

// FindAll busca todos os produtos
func (p *ProductDB) FindAll() ([]*entity.Product, error) {
	rows, err := p.DB.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	// Itera linha por linha do banco
	for rows.Next() {
		var id, name string
		var price float64
		// Scan mapeia as colunas para as variáveis
		if err := rows.Scan(&id, &name, &price); err != nil {
			return nil, err
		}
		// Recria a entidade. Note que aqui não validamos de novo (NewProduct),
		// pois confiamos que o que está no banco já é válido.
		products = append(products, &entity.Product{
			ID:    id,
			Name:  name,
			Price: price,
		})
	}
	return products, nil
}