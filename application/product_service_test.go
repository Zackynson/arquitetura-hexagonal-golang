package application_test

import (
	"testing"

	"github.com/Zackynson/go-hexagonal/application"
	mock_application "github.com/Zackynson/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeProduct := mock_application.NewMockProductInterface(ctrl)
	require.NotNil(t, fakeProduct)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	require.NotNil(t, persistence)

	persistence.EXPECT().Get(gomock.Any()).Return(fakeProduct, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("id")
	require.Nil(t, err)
	require.Equal(t, fakeProduct, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeProduct := mock_application.NewMockProductInterface(ctrl)
	require.NotNil(t, fakeProduct)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	require.NotNil(t, persistence)

	persistence.EXPECT().Save(gomock.Any()).Return(fakeProduct, nil).AnyTimes()
	require.NotNil(t, persistence)

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("any_name", 10)
	require.Nil(t, err)
	require.Equal(t, fakeProduct, result)

}

func TestProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeProduct := mock_application.NewMockProductInterface(ctrl)
	require.NotNil(t, fakeProduct)
	fakeProduct.EXPECT().Enable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	require.NotNil(t, persistence)

	persistence.EXPECT().Save(gomock.Any()).Return(fakeProduct, nil).AnyTimes()
	require.NotNil(t, persistence)

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(fakeProduct)
	require.Nil(t, err)
	require.Equal(t, fakeProduct, result)

}

func TestProductService_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeProduct := mock_application.NewMockProductInterface(ctrl)
	require.NotNil(t, fakeProduct)
	fakeProduct.EXPECT().Disable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	require.NotNil(t, persistence)

	persistence.EXPECT().Save(gomock.Any()).Return(fakeProduct, nil).AnyTimes()
	require.NotNil(t, persistence)

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Disable(fakeProduct)
	require.Nil(t, err)
	require.Equal(t, fakeProduct, result)

}
