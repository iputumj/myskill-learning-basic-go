package main

import "fmt"

func main() {
	point := 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Println("tidak lulus, nilai anda :", point)
	}
}
