package main

import "fmt"

func main() {
	/**
	 * map is an associative data type that is in the form of a value pair.
	 * For each data (or value) that is stored, a key is also prepared.
	 * The key must be unique, because it is used as a marker (or identifier) for accessing the corresponding value.
	 */
	var chicken = map[string]int{}
	fmt.Println(chicken) // map[]
	chicken["january"] = 50
	chicken["february"] = 40
	fmt.Println(chicken["january"]) // 50
	fmt.Println(chicken["maret"])   // 0
}
