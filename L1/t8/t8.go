package main

import (
	"fmt"
	"errors"
	"strconv"
	"math"
)

// установить для i-го бита числа num значени val
func SetBit(num * int64, i uint8, val uint8) error {
	// обработка ошибки
	if i > 64 || val > 1 {
		return errors.New("incorrect parametrs")
	} else {
		// создание маски из всех "1"
		mask := int64(math.MaxInt64 | math.MinInt64)
		// "0" на месте нужного бита
		mask -= int64(1 << i)
		// оставляем все биты, кроме i-го бита
		*num &= mask
		// ставим бит val 
		*num |= int64(val << i)
	
		return nil
	}
}

func PrintBits(a int64) {
	fmt.Println(strconv.FormatInt(a, 2))
}

func main() {

	a := int64(100)
	
	// до выставки бита
	fmt.Println("Before")
	PrintBits(a)

	// установка 1 на 3 бит
	SetBit(&a, 3, 1)

	// после
	fmt.Println("After")
	PrintBits(a)

	// проверка обработки ошибки
	fmt.Println("Error check")
	err := SetBit(&a, 3, 4)
	fmt.Println(err)
}