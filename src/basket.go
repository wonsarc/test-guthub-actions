package src

import (
	"fmt"
)

type BasketInterface interface {
	// AddProduct - Добавление товара в корзину в количестве count
	AddProduct(product Product, count int) error
	// DeleteProduct - Удаление товара из корзины
	DeleteProduct(id int) error
	// ListProducts - Список добавленных товаров
	ListProducts() []Product
	// GetPrice - Итоговая цена на все товары
	GetPrice() int
	// GetShippingCost - Цена доставки
	GetShippingCost() int
}

type Basket struct {
	Products []Product
}

const (
	maxCountBasket  = 30
	maxWeightBasket = 100
	minPriceProduct = 1
)

func (b *Basket) AddProduct(product Product, count int) error {
	currentWeight := 0
	currentCount := len(b.Products)

	for _, value := range b.Products {
		currentWeight += value.Weight
	}

	switch true {
	case count < 1:
		return nil
	case currentCount+count > maxCountBasket:
		return fmt.Errorf(ErrExceedMaxCount.Error(), maxCountBasket, currentCount)
	case currentWeight+product.Weight*count > maxWeightBasket:
		return fmt.Errorf(ErrExceedMaxWeight.Error(), maxWeightBasket, currentWeight)
	case product.Price < minPriceProduct:
		return fmt.Errorf(ErrInvalidProductPrice.Error(), minPriceProduct, product.Price)
	}

	for i := 0; i < count; i++ {
		b.Products = append(b.Products, product)
	}

	return nil
}

func (b *Basket) DeleteProduct(id int) error {
	newProducts := []Product{}
	isDelete := false

	for _, value := range b.Products {
		if value.Id == id {
			isDelete = true
		} else {
			newProducts = append(newProducts, value)
		}
	}

	if isDelete {
		b.Products = newProducts
		return nil
	}

	return fmt.Errorf(ErrProductNotFound.Error(), id)
}

func (b *Basket) GetPrice() int {
	priceProducts := getPriceProducts(b)
	priceDelivery := b.GetShippingCost()

	return priceDelivery + priceProducts
}

func (b *Basket) GetShippingCost() int {
	price := 0

	for _, value := range b.Products {
		price += value.Price
	}

	switch true {
	case price < 500:
		return 250
	case price < 1000:
		return 100
	default:
		return 0
	}
}

func (b *Basket) ListProducts() []Product {
	products := []Product{}
	products = append(products, b.Products...)

	return products
}

func getPriceProducts(b *Basket) int {
	price := 0

	for _, value := range b.Products {
		price += value.Price
	}

	return price
}

func NewBasket() BasketInterface {
	return &Basket{
		Products: []Product{},
	}
}
