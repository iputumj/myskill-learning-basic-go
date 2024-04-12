package main

import "fmt"

/**
 * Defer is used to terminate the execution of a statement just before the function block completes.
 * defer is commonly used to close connections to databases, etc.
 */

func main() {
	defer fmt.Println("Hello") // this line of code will be executed before the main block finishes
	fmt.Println("World")
}
