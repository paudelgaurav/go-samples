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

func RemoveDuplicate[K comparable](s []K) (result []K) {
	inResultMap := make(map[K]bool)

	for _, iter := range s {
		if _, ok := inResultMap[iter]; !ok {
			inResultMap[iter] = true
			result = append(result, iter)
		}
	}
	return result
}

func main() {
	fmt.Println(`Removing duplicate from ["a", "a", "b", "b", "c", "d", "d"]`)
	fmt.Println(RemoveDuplicate([]string{"a", "a", "b", "b", "c", "d", "d"}))
	fmt.Println(RemoveDuplicate([]int64{1, 1, 1, 2, 3, 4, 5, 4, 5}))
}
