package main

import (
	"fmt"
)

func removeElem(a []int, index int) []int {
	// через добавление к слайсу до индекса и все элементы после индекса
	return append(a[:index], a[index+1:]...)
	// можно через copy
	// copy(a[index:], a[index+1:])
	// return a[:len(a) - 1]
}

func main() {
	data := []int{0, 11, 22, 33, 44, 55, 66}

	fmt.Println(removeElem(data, 2))
}
