package main

import (
	"fmt"
)

func main() {

	// исходные данные
	data := []string{"cat", "cat", "cat", "dog", "tree"}
	// множество
	set := make(map[string]struct{})
	// создаем множество
	for _, v := range data {
		set[v] = struct{}{}
	}

	fmt.Println(set)
}