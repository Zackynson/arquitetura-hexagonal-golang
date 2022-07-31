package cli_test

import (
	"fmt"
	"testing"

	"github.com/Zackynson/go-hexagonal/adapters/cli"
	"github.com/Zackynson/go-hexagonal/application"
	mock_application "github.com/Zackynson/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_RUN(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := "ABC"
	productName := "product test"
	productPrice := 25.99
	productStatus := application.ENABLED

	productMock := mock_application.NewMockProductInterface(ctrl)

	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	serviceMock := mock_application.NewMockProductServiceInterface(ctrl)
	serviceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	// create
	result, err := cli.Run(serviceMock, "create", "", productName, productPrice)
	require.Nil(t, err)

	resultExpected := fmt.Sprintf("Product %s created successfully with name %s, price %f, and status %s", productId, productName, productPrice, productStatus)
	require.Equal(t, result, resultExpected)

	// enable
	result, err = cli.Run(serviceMock, "enable", productId, "0", 0)
	require.Nil(t, err)

	resultExpected = fmt.Sprintf("Product %s enabled successfully", productName)
	require.Equal(t, result, resultExpected)

	// disable
	result, err = cli.Run(serviceMock, "disable", productId, "0", 0)
	require.Nil(t, err)

	resultExpected = fmt.Sprintf("Product %s disabled successfully", productName)
	require.Equal(t, result, resultExpected)

	// get
	result, err = cli.Run(serviceMock, "get", productId, "0", 0)
	require.Nil(t, err)

	resultExpected = fmt.Sprintf("ProductID: %s\n Name: %s\nPrice: %f\nStatus: %s", productId, productName, productPrice, productStatus)
	require.Equal(t, result, resultExpected)

	// unsupported
	result, err = cli.Run(serviceMock, "", "", "0", 0)
	require.EqualError(t, err, "unsupported product action")
	require.Equal(t, "", result)

}
