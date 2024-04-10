package main

import "fmt"

func main() {
	/**
	 * An array is a collection of data of the same type, stored in a variable.
	 * arrays have a capacity whose value is determined at the time of creation,
	 * making the number of elements / data stored in the array cannot exceed what has been allocated.
	 */

	fruits := [4]string{"apple", "orange", "banana", "melon"}

	fmt.Println("length of array :", len(fruits))
	fmt.Println(fruits[1])
	fmt.Println(fruits)
}
