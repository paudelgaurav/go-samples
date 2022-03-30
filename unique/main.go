package main

import "fmt"

func unique(s []string) (result []string) {
	inResultMap := make(map[string]struct{})

	for _, str := range s {
		if _, found := inResultMap[str]; !found {
			inResultMap[str] = struct{}{}
			result = append(result, str)
		}
	}
	return result
}

func RemoveDuplicate[K comparable](s []K) (result []K) {
	inResultMap := make(map[K]struct{})

	for _, iter := range s {
		if _, found := inResultMap[iter]; !found {
			inResultMap[iter] = struct{}{}
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
