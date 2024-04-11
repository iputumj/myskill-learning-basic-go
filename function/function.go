package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"Michael", "Jordan"}
	printMessage("Hello", names)

}

func printMessage(message string, name []string) {
	fmt.Println(name)                     // [Michael Jordan]
	nameString := strings.Join(name, " ") // to convert an array to string
	fmt.Println(message, name)            // Hello [Michael Jordan]
	fmt.Println(message, nameString)      // Hello Michael Jordan
}
