package src

import (
	"fmt"
	"test-guthub-actions/src"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddProductInEmptyBasket(t *testing.T) {
	testCases := []struct {
		name    string
		product src.Product
		count   int
	}{
		{
			name:    "AddSuccessOneProduct",
			product: src.Product{Id: 1, Name: "Картошка", Weight: 5, Price: 15},
			count:   1,
		},
		{
			name:    "AddSuccess30Product",
			product: src.Product{Id: 1, Name: "Картошка", Weight: 3, Price: 15},
			count:   30,
		},
		{
			name:    "AddSuccess0Product",
			product: src.Product{Id: 1, Name: "Картошка", Weight: 5, Price: 15},
			count:   0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			basket := src.NewBasket()
			result := basket.AddProduct(tc.product, tc.count)
			assert.NoError(t, result)
			assert.Equal(t, tc.count, len(basket.ListProducts()))
		})
	}
}

func TestErrorsAddProduct(t *testing.T) {

	testCases := []struct {
		name          string
		product       src.Product
		count         int
		expectedError error
	}{
		{
			name:          "AddSuccesOneProduct",
			product:       src.Product{Id: 1, Name: "Картошка", Weight: 1, Price: 15},
			count:         2,
			expectedError: fmt.Errorf(src.ErrExceedMaxCount.Error(), 30, 29),
		},
		{
			name:          "AddSucces30Product",
			product:       src.Product{Id: 1, Name: "Картошка", Weight: 14, Price: 15},
			count:         1,
			expectedError: fmt.Errorf(src.ErrExceedMaxWeight.Error(), 100, 87),
		},
		{
			name:          "AddSucces0Product",
			product:       src.Product{Id: 1, Name: "Картошка", Weight: 1, Price: 0},
			count:         1,
			expectedError: fmt.Errorf(src.ErrInvalidProductPrice.Error(), 1, 0),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			basket := src.NewBasket()
			product := src.Product{Id: 2, Name: "Капуста", Weight: 3, Price: 20}
			basket.AddProduct(product, 29)
			result := basket.AddProduct(tc.product, tc.count)
			assert.Equal(t, tc.expectedError.Error(), result.Error())
		})
	}
}

func TestDeleteProduct(t *testing.T) {

	testCases := []struct {
		name          string
		id            int
		expectedError error
	}{
		{
			name:          "SuccessDelete",
			id:            1,
			expectedError: nil,
		},
		{
			name:          "ErrorDelete",
			id:            2,
			expectedError: fmt.Errorf(src.ErrProductNotFound.Error(), 2),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			basket := src.NewBasket()
			basket.AddProduct(src.Product{Id: 1, Name: "Картошка", Weight: 1, Price: 2}, 1)
			result := basket.DeleteProduct(tc.id)

			if tc.expectedError == nil {
				assert.NoError(t, result)
				assert.Equal(t, 0, len(basket.ListProducts()))

			} else {
				assert.Equal(t, tc.expectedError.Error(), result.Error())
				assert.Equal(t, 1, len(basket.ListProducts()))
			}

		})
	}
}

func TestListProducts(t *testing.T) {
	testCases := []struct {
		name           string
		product        src.Product
		count          int
		expectedResult []src.Product
	}{
		{
			name:           "GetNotEmptyList",
			product:        src.Product{Id: 1, Name: "Картошка", Weight: 14, Price: 15},
			count:          1,
			expectedResult: []src.Product{{Id: 1, Name: "Картошка", Weight: 14, Price: 15}},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			basket := src.NewBasket()
			basket.AddProduct(tc.product, tc.count)
			result := basket.ListProducts()

			assert.Equal(t, tc.count, len(result))
			assert.Equal(t, tc.product.Id, result[0].Id)
			assert.Equal(t, tc.product.Name, result[0].Name)
			assert.Equal(t, tc.product.Price, result[0].Price)
			assert.Equal(t, tc.product.Weight, result[0].Weight)

		})
	}
}

func TestEmptyListProducts(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult []src.Product
	}{
		{
			name:           "GetNotEmptyList",
			expectedResult: []src.Product{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			basket := src.NewBasket()
			result := basket.ListProducts()

			assert.Equal(t, 0, len(result))
		})
	}
}

func TestGetPrices(t *testing.T) {
	testCases := []struct {
		name           string
		product        src.Product
		count          int
		expectedResult int
	}{
		{
			name:           "More1Count",
			product:        src.Product{Id: 1, Name: "Картошка", Weight: 14, Price: 15},
			count:          5,
			expectedResult: 325,
		},

		{
			name:           "1ProductPrice",
			product:        src.Product{Id: 1, Name: "Картошка", Weight: 14, Price: 15},
			count:          1,
			expectedResult: 265,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			basket := src.NewBasket()
			basket.AddProduct(tc.product, tc.count)
			result := basket.GetPrice()

			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestGetEmptyPrices(t *testing.T) {
	t.Run("EmptiListPrice", func(t *testing.T) {
		basket := src.NewBasket()
		result := basket.GetPrice()
		assert.Equal(t, 250, result)
	})
}

func TestGetShippingCost(t *testing.T) {
	testCases := []struct {
		name           string
		product        src.Product
		expectedResult int
	}{
		{
			name:           "<500",
			product:        src.Product{Id: 1, Name: "Картошка", Weight: 14, Price: 499},
			expectedResult: 250,
		},

		{
			name:           "<1000Min",
			product:        src.Product{Id: 1, Name: "Картошка", Weight: 14, Price: 500},
			expectedResult: 100,
		},
		{
			name:           "<1000Max",
			product:        src.Product{Id: 1, Name: "Картошка", Weight: 14, Price: 999},
			expectedResult: 100,
		},
		{
			name:           ">=1000",
			product:        src.Product{Id: 1, Name: "Картошка", Weight: 14, Price: 1000},
			expectedResult: 0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			basket := src.NewBasket()
			basket.AddProduct(tc.product, 1)
			result := basket.GetShippingCost()

			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
