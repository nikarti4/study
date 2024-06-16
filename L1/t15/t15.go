// предложенный пример приводит к выходу за границы слайса

package main

import (
	"fmt"
	"strconv"
)

var justString string

func createHugeString(a int) string {
	return strconv.Itoa(a)
}

func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Println(v)
	justString = v[:]
}

func main() {
	someFunc()
}
