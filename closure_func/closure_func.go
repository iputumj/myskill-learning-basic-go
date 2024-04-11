package main

import "fmt"

/**
 * A closure definition is a function that can be stored in a variable.
 * By applying this concept, we can create a function inside a function, or even create a function that returns a function
 */

func main() {
	getMinMax := func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	numbers := []int{2, 3, 4, 5, 3, 2}
	min, max := getMinMax(numbers)
	fmt.Println("data: ", numbers)
	fmt.Println("min: ", min)
	fmt.Println("max: ", max)
}
