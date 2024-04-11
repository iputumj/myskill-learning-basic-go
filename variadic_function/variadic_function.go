package main

import "fmt"

/**
 * Go adopts the concept of variadic function or function creation with unlimited similar parameters,
 * the meaning of unlimited is that the number of parameters inserted when calling the function can be any.
 */
func main() {
	avg := calculate(2, 4, 5, 6, 7, 8, 8, 3, 4, 5)
	fmt.Println("Average :", avg)
}

func calculate(numbers ...int) float64 {
	total := 0

	for _, number := range numbers {
		total += number
	}

	avg := float64(total) / float64(len(numbers))
	return avg

}
