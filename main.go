package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"test-guthub-actions/src"
)

func main() {
	basket := src.NewBasket()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите команду (add, list, price, exit):")
		scanner.Scan()
		command := scanner.Text()

		switch strings.ToLower(command) {
		case "add":
			var id, count int
			var name string
			var price, weight int

			fmt.Println("Введите ID продукта:")
			fmt.Scan(&id)
			fmt.Println("Введите имя продукта:")
			fmt.Scan(&name)
			fmt.Println("Введите цену продукта:")
			fmt.Scan(&price)
			fmt.Println("Введите вес продукта:")
			fmt.Scan(&weight)
			fmt.Println("Введите количество:")
			fmt.Scan(&count)

			product := src.NewProduct(id, name, price, weight)
			err := basket.AddProduct(*product.(*src.Product), count)
			if err != nil {
				fmt.Println("Ошибка при добавлении продукта:", err)
			} else {
				fmt.Println("Продукт добавлен в корзину.")
			}

		case "list":
			products := basket.ListProducts()
			fmt.Println("Продукты в корзине:")
			for _, p := range products {
				fmt.Printf("ID: %d, Name: %s, Price: %d, Weight: %d\n", p.GetID(), p.GetName(), p.GetPrice(), p.GetWeight())
			}

		case "price":
			totalPrice := basket.GetPrice()
			shipPrice := basket.GetShippingCost()
			productPrice := totalPrice - shipPrice
			fmt.Printf("Стоимость корзины: %d, Стоимость доставки: %d, Итоговая цена: %d\n", productPrice, shipPrice, totalPrice)
		case "exit":
			fmt.Println("Выход из программы.")
			return

		default:
			fmt.Println("Неизвестная команда. Попробуйте снова.")
		}
	}
}
