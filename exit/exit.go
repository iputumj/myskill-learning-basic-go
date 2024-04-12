package main

import (
	"fmt"
	"os"
)

/**
 * Exit digunakan untuk menghentikan program secara paksa (ingat, menghentikan program, tidak seperti return yang hanya menghentikan blok kode)
 */
func main() {
	defer fmt.Println("Hello")
	os.Exit(1)
	fmt.Println("World")
}
