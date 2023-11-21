package main

import (
	"fmt"
)

func isTriangle(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}

func main() {
	var a, b, c int
	fmt.Println("Введите число a;")
	fmt.Scan(&a)
	fmt.Println("Введите число b;")
	fmt.Scan(&b)
	fmt.Println("Введите число c;")
	fmt.Scan(&c)
	if isTriangle(a, b, c) {
		fmt.Println("Может получиться треугольник")
	} else {
		fmt.Println("Какая то шляпа :(")
	}
}
