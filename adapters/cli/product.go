package cli

import (
	"errors"
	"fmt"

	"github.com/Zackynson/go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64) (string, error) {

	var result = ""

	switch action {
	case "create":
		p, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s created successfully with name %s, price %f, and status %s", p.GetId(), p.GetName(), p.GetPrice(), p.GetStatus())
		return result, nil

	case "enable":
		p, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		res, err := service.Enable(p)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s enabled successfully", res.GetName())
		return result, nil

	case "disable":
		p, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		res, err := service.Disable(p)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s disabled successfully", res.GetName())
		return result, nil

	case "get":
		p, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("ProductID: %s\n Name: %s\nPrice: %f\nStatus: %s", p.GetId(), p.GetName(), p.GetPrice(), p.GetStatus())

		return result, nil

	default:
		return "", errors.New("unsupported product action")
	}
}
