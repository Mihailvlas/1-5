package main

import (
	"errors"
	"fmt"
)

func positiveNumbers(n int) ([]int, error) {
	if n <= 0 {
		return nil, errors.New("введенное значение должно быть положительным числом")
	}
	numbers := []int{}
	for i := 1; i <= n; i++ {
		numbers = append(numbers, i)
	}
	return numbers, nil
}

func main() {
	var maxValue int
	fmt.Print("Введите максимальное значение: ")
	_, err := fmt.Scan(&maxValue)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	numbers, err := positiveNumbers(maxValue)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Позитивные числа до", maxValue, ":", numbers)
}
