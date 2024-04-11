package main

import "fmt"

/**
 * Functions can be made into variable data type,
 * and it is possible to make it a parameter as well
 */

func main() {
	result := filter([]string{"this is", "data"}, func(each string) bool {
		return true
	})
	fmt.Println(result)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
