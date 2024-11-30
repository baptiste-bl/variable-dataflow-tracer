package main

import (
	"fmt"
	"math"
)

func CalculateArea(radiusTest2 float64) float64 {
	area := radiusTest2
	test = math.Pi * area * area
}

func DoubleArea(area float64) float64 {
	return 2 * area
}

func CalculateAndDouble(radiusTest float64) float64 {
	area := CalculateArea(radiusTest)
	test := DoubleArea(radiusTest)
	print(test)
	doubleArea := DoubleArea(area)
	return doubleArea
}

func test() {
	radius := 5.0
	result := CalculateAndDouble(radius)
	radius := 10.0
	fmt.Println(result)
}

func main() {
	test()
}
