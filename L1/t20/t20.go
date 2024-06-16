package main


import (
	"fmt"
	"strings"
	"slices"
)

func reverseWords(s string) string {
	
	var out string
	
	// разбиваем строку на слова
	words := strings.Fields(s)
	// меняем местами элементы получившегося слайса
	slices.Reverse(words)
	// формируем выходную строку
	for _, word := range words {
		out += (word + " ")
	}

	return out
}

func main() {
	fmt.Println(reverseWords("snow dog sun"))
	fmt.Println(reverseWords("Разработать программу, которая переворачивает слова в строке"))
}