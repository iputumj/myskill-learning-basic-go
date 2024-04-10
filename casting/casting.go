package main

import "fmt"

func main() {
	/**
	 * A casting type is a data type resulting from the conversion of another data type
	 */
	var a float64 = float64(24)
	fmt.Println(a)

	var b int32 = int32(24.00)
	fmt.Println(b)

	/**
	 * A method of declaring a variable whose data type is determined by the data type of its value, which is contradictory to the first method.
	 * With this type of method, the var keywork and data type don't need to be written.
	 */
}
