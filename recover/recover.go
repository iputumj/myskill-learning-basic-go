package main

import (
	"errors"
	"fmt"
	"strings"
)

/**
 * Recover is useful for handling panic errors
 * When a panic erorr occurs, recover will take over the goroutine that's panicking (the panic message will not appear)
 */
func main() {
	defer catch()

	var name string
	fmt.Println("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Hello", name)
	} else {
		panic(err.Error())
		fmt.Println("End")
	}
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	} else {
		return true, nil
	}
}
