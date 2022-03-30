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

func itemInList[T comparable](item T, itemList []T) bool {
	for _, i := range itemList {
		if i == item {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(`Checking if "a" exists in ["a", "e", "i"]`)
	fmt.Println(itemInList("a", []string{"a", "e", "i"}))
	fmt.Println(itemInList(404, []int64{200, 201, 202}))
}
