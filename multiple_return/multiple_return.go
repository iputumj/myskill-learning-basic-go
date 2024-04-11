package main

import (
	"fmt"
	"math"
)

/**
 * Generally, functions only have one return only
 * If there is a need for more data to be returned, it is common to use data types like map, slice or struct
 * But in Golang, it can have multiple return
 */
func main() {
	value := 5.00
	area, circumference := calculate(value)
	fmt.Println(area)
	fmt.Println(circumference)
}

func calculate(value float64) (float64, float64) {
	area := math.Pi * math.Pow(value/2, 2)
	circumference := math.Pi * value

	return area, circumference
}
