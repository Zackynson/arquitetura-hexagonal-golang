package db

import (
	"database/sql"

	"github.com/Zackynson/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{
		db: db,
	}
}

func (p *ProductDB) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := p.db.Prepare("select id, name, price, status from products where id = ?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDB) create(product application.ProductInterface) (application.ProductInterface, error) {
	insertQuery := `insert into products (id, name, price, status) values (?,?,?,?)`

	stmt, err := p.db.Prepare(insertQuery)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetId(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus())

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDB) update(product application.ProductInterface) (application.ProductInterface, error) {
	updateQuery := `update products set name = ?, status = ?, price = ? where id = ?`

	_, err := p.db.Exec(
		updateQuery,
		product.GetName(),
		product.GetPrice(),
		product.GetId(),
		product.GetStatus())

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDB) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	p.db.QueryRow(`select id from products where id = ?`, product.GetId()).Scan(&rows)

	if rows == 0 {
		insertedProduct, err := p.create(product)
		if err != nil {
			return nil, err
		}

		return insertedProduct, nil
	}

	updatedProduct, err := p.update(product)

	if err != nil {
		return nil, err
	}

	return updatedProduct, nil

}
