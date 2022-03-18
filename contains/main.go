package main

import (
	"fmt"
)

func contains(elems []string, value string) bool {
	for _, s := range elems {
		if value == s {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(`Checking if "a" exists in ["a", "e", "i"]`)
	fmt.Println(contains([]string{"a", "e", "i"}, "a"))
}
