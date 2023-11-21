package main

import (
	"fmt"
	"strings"
)

func toCamelCase(s string) string {
	words := strings.Split(s, "-")
	camelCase := ""
	for i, word := range words {
		if i == 0 {
			camelCase += strings.Title(word)
			camelCase += strings.Title(word)
		}
	}
	return camelCase
}

func main() {
	input := "hello-world"
	output := toCamelCase(input)
	fmt.Println(output)
}
