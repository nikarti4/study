package main

import (
	"fmt"
)

func quickSort(a []int) []int {

	// уже отсортирован
	if len(a) < 2 {
		return a
	}

	// правая (считаем опорным элементом) и левая границы
	l := 0
	r := len(a) - 1

	// проходимся по всем элементам, и элементы меньше
	// опорного помещаем влево, и сдвигаем левую границу
	for i := range a {
		if a[i] < a[r] {
			a[i], a[l] = a[l], a[i]
			l++
		}
	}

	// ставим опорный таким образом, чтобы
	// слева были числа меньше опорного, а справа больше
	a[l], a[r] = a[r], a[l]

	// сортируем подмассивы слева и справа от опорного
	quickSort(a[:l])
	quickSort(a[l+1:])

	return a
}

func main() {

	data := []int{5, -1, 3, 7, -100, 200, 4, 0, 3, 8}

	fmt.Println("before: ", data)
	fmt.Println("after: ", quickSort(data))
}
