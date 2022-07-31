package db_test

import (
	"database/sql"
	"log"
	"testing"

	db "github.com/Zackynson/go-hexagonal/adapters/db"
	"github.com/Zackynson/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) error {

	createTableQuery := `create table products("id" string, "name" string, "price" float, "status" string);`

	stmt, err := db.Prepare(createTableQuery)
	if err != nil {
		return err
	}

	defer stmt.Close()
	stmt.Exec()

	return nil
}

func createProduct(db *sql.DB) {
	insert := `insert into products (id, name, price, status) values (1, "teste", 10, "enabled")`
	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatalf("error creating product %v", err)
	}

	stmt.Exec()
	stmt.Close()
}

func TestProductDB_Get(t *testing.T) {
	setup()
	defer Db.Close()

	productDB := db.NewProductDB(Db)

	product, err := productDB.Get("1")
	require.Nil(t, err)
	require.Equal(t, "teste", product.GetName())
	require.Equal(t, "1", product.GetId())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, "enabled", product.GetStatus())

	product, err = productDB.Get("2")
	require.Nil(t, product)
	require.NotNil(t, err)
}

func TestProductDB_Save(t *testing.T) {
	setup()
	defer Db.Close()

	productDB := db.NewProductDB(Db)
	product := application.NewProduct()
	product.Name = "teste"
	product.Status = application.ENABLED
	product.Price = 10

	insertedProduct, err := productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetName(), insertedProduct.GetName())
	require.Equal(t, product.GetId(), insertedProduct.GetId())
	require.Equal(t, product.GetPrice(), insertedProduct.GetPrice())
	require.Equal(t, insertedProduct.GetStatus(), insertedProduct.GetStatus())

	product.Name = "test updated"

	insertedProduct2, err := productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, "test updated", insertedProduct2.GetName())
	require.Equal(t, insertedProduct2.GetId(), insertedProduct.GetId())
	require.Equal(t, insertedProduct2.GetPrice(), insertedProduct.GetPrice())
	require.Equal(t, insertedProduct2.GetStatus(), insertedProduct.GetStatus())

}
