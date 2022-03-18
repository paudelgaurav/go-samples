package main

import "fmt"

func unique(s []string) []string {
	inResultMap := make(map[string]bool)
	var result []string

	for _, str := range s {
		if _, ok := inResultMap[str]; !ok {
			inResultMap[str] = true
			result = append(result, str)
		}
	}
	return result
}

func main() {
	fmt.Println(`Removing duplicate from ["a", "a", "b", "b", "c", "d", "d"]`)
	fmt.Println(unique([]string{"a", "a", "b", "b", "c", "d", "d"}))
}
