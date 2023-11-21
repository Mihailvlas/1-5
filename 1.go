package main

import (
	"fmt"
)

func findCentury(year int) (int, error) {
	if year <= 0 {
		return 0, fmt.Errorf("введенный год должен быть положительным числом")
	}
	century := (year-1)/100 + 1
	return century, nil
}

func main() {
	var inputYear int
	fmt.Print("Введите год: ")
	_, err := fmt.Scan(&inputYear)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	century, err := findCentury(inputYear)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Год", inputYear, "находится в", century, "веке.")
}
