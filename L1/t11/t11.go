package main

import (
	"fmt"
)

func intersect(a1 []int, a2 []int) []int {
	// возвращаемое значение
	var out[]int
	// ключ - элемент множества
	m := make(map[int]struct{})

	// записываем все ключи
	for _, v := range a1 {
		m[v] = struct{}{}
	}

	// проверяем, есть ли элемент второго подмн. в ключах, 
	// если есть - добавляем в слайс
	for _, v := range a2 {
		_, ok := m[v]

		if ok {
			out = append(out, v)
		}
	}

	return out
}

func main() {

	// два неупорядоченных множества
	a1 := []int{1, 3, 3, 5, 9, 10, -8, 7}
	a2 := []int{2, 0, 1, 11, 3, -4, -8, 4, 6, 22}

	fmt.Println("a1: ", a1)
	fmt.Println("a2: ", a2)

	out := intersect(a1, a2)

	fmt.Println("out: ", out)
}