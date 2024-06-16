package main

import (
	"fmt"
)

func main() {
	// инициализируем две переменных
	a := 12
	b := 34

	fmt.Println("before", a, b)

	// математика
	a = a + b
	b = a - b
	a = a - b

	fmt.Println("after", a, b)
}
