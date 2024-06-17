package main

import (
	"t24/point"
	"math"
	"fmt"
)

// расстояние от точки p2 до p1
func distance(p1, p2 point.Point) float64 {
	
	t1 := math.Pow(p2.GetX() - p1.GetX(), 2)
	t2 := math.Pow(p2.GetY() - p1.GetY(), 2)
	
	return math.Sqrt(t1 + t2)
}

func main() {
	p1 := point.NewPoint(1, 2)
	p2 := point.NewPoint(3, 4)

	fmt.Println("Distance between", p1, " and ", p2, " = ", distance(p1, p2))
}