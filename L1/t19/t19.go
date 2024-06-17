package main

import (
	"fmt"
)

func reverse(s string) string {
	var out string

	// выделяем место под слайс рун
	temp := make([]rune, len(s))

	// при использовании for range
	// автопреобразование в rune
	for i, r := range s {
		temp[len(s)-i-1] = r
	}

	// обратная конвертация в строку
	out = string(temp)

	return out
}

func main() {
	fmt.Println(reverse("Hello"))
	fmt.Println(reverse("главрыба"))
	fmt.Println(reverse("abc😀defg"))
}
