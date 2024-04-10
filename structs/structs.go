package main

import "fmt"

/**
 * structs are a collection of variable (or property) and or function (or method) definitions, wrapped as a new data type with a specific name.
 * The properties inside a struct, the data type can vary. Similar to a map, except that the key is defined at the beginning, and the data type of each item can be different.
 */
type student struct {
	name  string
	grade int
}

func main() {
	var s1 student
	s1.name = "Haerin"
	s1.grade = 2

	fmt.Println(s1)
	fmt.Println(s1.name)
	fmt.Println(s1.grade)
}
