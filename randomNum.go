package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Установка начального значения для генератора случайных чисел

	var a, b int
	b = rand.Intn(10) // Генерируем случайное число один раз,т.е. не генерерует постоянно после каждого ввода

	fmt.Println("Угадайте число от 0 до 10")
	for {
		fmt.Print("Введите число: ")
		fmt.Scan(&a)

		if a == b {
			fmt.Println("Верно!")
			break // Выход из цикла при правильном ответе
		} else {
			fmt.Println("Неверно, попробуйте снова.")
		}
	}
}
