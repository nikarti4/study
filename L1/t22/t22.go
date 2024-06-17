package main

import (
	"fmt"
	"math/big"
	"math"
)

// большие числа можно инициализоравть через
// a := big.NewInt(число) (если число влезает в int64)
// или
// a := new(big.Int)
// a.SetString("число", 10)

func main() {
	// создаем большие переменные
	a := big.NewFloat(math.Pow(2, 30))
	b := big.NewFloat(math.Pow(2, 28))

	// создаем результаты операций
	resAdd := big.NewFloat(0)
	resSub := big.NewFloat(0)
	resMul := big.NewFloat(0)
	resDiv := big.NewFloat(0)

	// вызываем методы операций
	resAdd.Add(a, b)
	resSub.Sub(a, b)
	resMul.Mul(a, b)
	resDiv.Quo(a, b)

	// вывод результатов
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("resAdd = ", resAdd)
	fmt.Println("resSub = ", resSub)
	fmt.Println("resMul = ", resMul)
	fmt.Println("resDiv = ", resDiv)
}