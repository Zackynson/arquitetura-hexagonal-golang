package application_test

import (
	"testing"

	"github.com/Zackynson/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func genUuid() string {
	return uuid.NewV4().String()
}

func TestProduct_Enable(t *testing.T) {
	product := application.Product{
		Name:   "Product Test",
		Id:     genUuid(),
		Status: application.DISABLED,
		Price:  10,
	}

	err := product.Enable()
	require.Nil(t, err)

	product2 := application.Product{
		Name:   "Product Test",
		Id:     genUuid(),
		Status: application.DISABLED,
		Price:  0,
	}

	err = product2.Enable()
	require.EqualError(t, err, "product price must be greater than zero to be enabled it")
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{
		Name:   "Product Test",
		Id:     genUuid(),
		Status: application.ENABLED,
		Price:  0,
	}

	err := product.Disable()
	require.Nil(t, err)

	product2 := application.Product{
		Name:   "Product Test",
		Id:     genUuid(),
		Status: application.ENABLED,
		Price:  10,
	}

	err = product2.Disable()
	require.EqualError(t, err, "product price must be zero to be disabled it")
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{
		Name:   "Product Test",
		Id:     genUuid(),
		Price:  10,
		Status: application.ENABLED,
	}

	// must pass validation
	isValid, err := product.IsValid()
	require.Nil(t, err)
	require.Equal(t, true, isValid)

	// invalid status
	product.Status = "INVALID"
	isValid, err = product.IsValid()
	require.EqualError(t, err, "Product status must be enabled or disabled")
	require.Equal(t, false, isValid)

	// valid status
	product.Status = application.ENABLED
	isValid, err = product.IsValid()
	require.Nil(t, err)
	require.Equal(t, true, isValid)

	// invalid price
	product.Price = -10
	isValid, err = product.IsValid()
	require.EqualError(t, err, "Product price must be greater or equal zero")
	require.Equal(t, false, isValid)

	// valid price
	product.Price = 0
	isValid, err = product.IsValid()
	require.Nil(t, err)
	require.Equal(t, true, isValid)

	// invalid name
	product.Name = ""
	isValid, err = product.IsValid()
	require.EqualError(t, err, "Name: non zero value required")
	require.Equal(t, false, isValid)

	// valid name
	product.Name = "any name"
	isValid, err = product.IsValid()
	require.Nil(t, err)
	require.Equal(t, true, isValid)

	// invalid id
	product.Id = "non_uuid_id"
	isValid, err = product.IsValid()
	require.EqualError(t, err, "Id: non_uuid_id does not validate as uuidv4")
	require.Equal(t, false, isValid)

	// valid id
	product.Id = genUuid()
	isValid, err = product.IsValid()
	require.Nil(t, err)
	require.Equal(t, true, isValid)
}

func TestProduct_GetId(t *testing.T) {
	productId := genUuid()

	product := application.Product{
		Name:   "Product Test",
		Id:     productId,
		Status: application.ENABLED,
		Price:  0,
	}

	id := product.GetId()
	require.Equal(t, id, product.Id)
}

func TestProduct_GetName(t *testing.T) {
	productId := genUuid()

	product := application.Product{
		Name:   "Product Test",
		Id:     productId,
		Status: application.ENABLED,
		Price:  0,
	}

	name := product.GetName()
	require.Equal(t, name, product.Name)
}

func TestProduct_GetStatus(t *testing.T) {
	productId := genUuid()

	product := application.Product{
		Name:   "Product Test",
		Id:     productId,
		Status: application.ENABLED,
		Price:  0,
	}

	status := product.GetStatus()
	require.Equal(t, status, product.Status)
}

func TestProduct_GetPrice(t *testing.T) {
	productId := genUuid()

	product := application.Product{
		Name:   "Product Test",
		Id:     productId,
		Status: application.ENABLED,
		Price:  0,
	}

	price := product.GetPrice()
	require.Equal(t, price, product.Price)
}
