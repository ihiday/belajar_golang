package main

import (
	"fmt"
)

func main() {
	str := "selamat malam"
	var chars []string

	for _, char := range str {
		chars = append(chars, string(char))
	}

	fmt.Println(chars)
}
