package main

import (
	"fmt"
	"strconv"
)

/**
 * Error is a data type.
 * It has 1 property, the Error() method, which returns the details of error message in string.
 * Error is a data type whose content can be nil.
 */
func main() {
	var input string
	var number int
	var err error

	fmt.Println("Type some number: ")
	fmt.Scanln(&input)

	number, err = strconv.Atoi(input)
	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}

}
