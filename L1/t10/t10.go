package main

import (
	"fmt"
	"math"
)

// округление до ближайшего числа, кратного 10
func roundTo10th(inp float64) float64 {
	var out float64

	// если число отрицательное - округлить до большего
	// если отрицательное - до меньшего
	if inp < 0 {
		out = math.Ceil(inp/10) * 10
	} else {
		out = math.Floor(inp/10) * 10
	}

	return out
}

func main() {
	// исходные данные
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// подмножества по ключам с шагом 10
	set := make(map[int][]float64)

	for _, v := range data {
		// выделяем целую часть
		i, _ := math.Modf(v)
		// округляем до ближ. с шагом 10
		i_t := int(roundTo10th(i))
		// добавляем
		set[i_t] = append(set[i_t], v)
	}

	fmt.Println(set)
}
