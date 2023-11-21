package main

import (
	"fmt"
	"strings"
)

func accumulateString(str string) string {
	result := ""
	for i, char := range str {
		if i > 0 {
			result += "-"
		}
		result += strings.Repeat(string(char), i+1)
	}
	return result
}

func main() {
	var str string
	fmt.Print("Введите строку: ")
	fmt.Scan(&str)
	result := accumulateString(str)
	fmt.Println(result)
}
