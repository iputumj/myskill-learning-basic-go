package main

import (
	"errors"
	"fmt"
	"strings"
)

/**
 * Panic is used to display the stack trace error while stopping the goroutine flow.
 * (Since main() is also a goroutine, the same behavior applies).
 *
 * After a panic, the process will be stopped, anything after is not executed except the process that has been deferred before (will appear before the panic error).
 */
func main() {
	var name string
	fmt.Println("Type your name: ")
	fmt.Scanln(&name)
	if valid, err := validate(name); valid {
		fmt.Println("Hello", name)
	} else {
		defer fmt.Println("End 1")
		panic(err.Error())
		fmt.Println("End 2")
	}
}

func validate(name string) (bool, error) {
	if strings.TrimSpace(name) == "" {
		return false, errors.New("name cannot be empty")
	}
	return true, nil
}
