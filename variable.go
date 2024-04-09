package main

import "fmt"

func main() {
	/**
	 * A variable is a container for storing data, a variable has a data type, which is the type of data it holds.
	 */
	var firstName string = "my"
	var lastName string = "skill"
	city := "Bali"

	fmt.Println("Hello", firstName, lastName)
	fmt.Println(city)

	/**
	 * constants are similar to variables in that they are data storage containers, but the data cannot be changed after being assigned
	 */
	const country = "Indonesia"
	// country = "Korea" //error

	fmt.Println(country)
}
