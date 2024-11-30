package src

import "errors"

var (
	ErrExceedMaxCount      = errors.New("превышен лимит на количество товаров в корзине. Лимит: %d. Текущее количество товаров в корзине: %d")
	ErrExceedMaxWeight     = errors.New("превышен лимит на вес товаров в корзине. Лимит: %d. Текущий вес товаров в корзине: %d")
	ErrInvalidProductPrice = errors.New("указана неверная цена. Минимальная цена: %d. Текущая цена: %d")
	ErrProductNotFound     = errors.New("не найден товар с id: %d")
)
