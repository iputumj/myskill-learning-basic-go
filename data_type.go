package main

import "fmt"

func main() {
	/**
	 * non-decimal numeric
	 */
	var positiveNum uint8 = 89
	var negativeNum = -1243423644

	fmt.Println("Positive Number : ", positiveNum)
	fmt.Println("Nagative Number : ", negativeNum)

	/**
	 * Decimal numeric
	 */
	var decimalNum = 2.0
	fmt.Println("Decimal Number : ", decimalNum)

	/**
	 * Boolean
	 */
	var isBoolean = true
	fmt.Println(isBoolean)

	/**
	 * String
	 */
	var message = "Hello"
	fmt.Println(message)

	/**
	 * nil value, zero value || default value
	 *
	 * default value of string is ""
	 * default value of bool is false
	 * default value of non-decimal numeric is 0
	 * default value of decimal numeric is 0.0
	 */
}
