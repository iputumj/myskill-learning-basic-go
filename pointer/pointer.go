package main

import "fmt"

func main() {
	/**
	 * Pointers are memory references or addresses.
	 * Variable pointer means a variable that contains the memory address of a value.
	 * For example, a variable with integer data type has a value of 4, so what is meant by pointer is the memory address where the value of 4 is stored, not the value of 4 itself.
	 *
	 * Pointers are useful for transferring data that has a large capacity, through a function. pointers are closely related to arrays, so pointers can replace the function of array variables.
	 *
	 * there are two important things to know:
	 * 1. Ordinary variables can have their pointer values retrieved, by adding an ampersand (&) right before the variable name. this method is called referencing.
	 * 2. And vice versa, the original value of a pointer variable can also be retrieved, by adding an asterisk (*) right before the variable name. This method is called deferencing.
	 */
	var numberA int = 4
	var numberB *int = &numberA

	// if u try to re-assign the value of pointer numberB, then the value of number A will also change
	*numberB = 8

	// the result is 0x14000110018. this is memory of numberA
	fmt.Println(numberB)

	// if u want to get the value, use *
	fmt.Println(*numberB)
	fmt.Println(numberA)
}
