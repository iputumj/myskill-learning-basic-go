package main

import (
	"errors"
	"fmt"
	"strings"
)

/**
 * Besides utilizing the error return of an internal function, we can also create our own error object using the errors function.
 * New() (must import the error package first)
 */

func main() {
	var name string
	fmt.Println("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Hello", name)
	} else {
		fmt.Println(err.Error())
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Cannot be empty")
	} else {
		return true, nil
	}
}
