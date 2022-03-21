package main

import "fmt"

func divide(a int, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovering from panic in divide function error is: %v \n", r)
		}
	}()

	return a / b
}

func main() {
	x := 20
	y := 10
	z := 0
	fmt.Println(divide(x, y))
	fmt.Print(divide(x, z))
}
