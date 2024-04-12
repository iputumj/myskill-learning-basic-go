package main

import (
	"fmt"
	"package/lib"
)

/**
 * Go is designed to be a language that encourages good software engineering practices.
 * An important part of high-quality software is code reuse - embodied in the "Don't repeat yourself" principle.
 * This Go package is a piece of code that is broken down and simplified for reuse.
 */
func main() {
	xs := []float64{1, 2, 3, 4}
	avg := lib.Average(xs)
	fmt.Println(avg)
}
