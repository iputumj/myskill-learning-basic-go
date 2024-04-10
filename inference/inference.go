package main

import "fmt"

func main() {
	/**
	 * A method of declaring a variable whose data type is determined by the data type of its value, which is contradictory to the first method.
	 * With this type of method, the var keywork and data type don't need to be written.
	 */
	var city string = "Jakarta"
	city2 := "Bandung"
	fmt.Println(city)
	fmt.Println(city2)
}
