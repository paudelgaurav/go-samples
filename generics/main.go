package main

import "fmt"

func SumInts(m map[string]int64) (s int64) {
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) (s float64) {
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) (s V) {
	for _, v := range m{
		s += v
	}
	return s
}

func main() {
	ints := map[string]int64{
		"first":  12,
		"second": 13,
	}

	floats := map[string]float64{
		"first":  12.33,
		"second": 13.41,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))
}