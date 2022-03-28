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
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	int8 | int64 | float64
}

func SumNumbers[N Number](nums []N) (sum N) {
	for _, v := range nums {
		sum += v
	}
	return sum
}

// any can be used instead of interface{}
func Print[T any] (s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

type CustomSlice[T Number] []T

func PrintCustomSliceElems[N Number, T CustomSlice[N]] (s T) {
	for _, v := range s {
		fmt.Println(v)
	}
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

	int_array := []int64{32, 64, 96, 128}
	float_array := []float64{32.0, 64.0, 96.1, 128.2}
	bytes_array := []int8{8, 16, 24, 32}

	fmt.Println(SumNumbers(int_array))
	fmt.Println(SumNumbers(float_array))
	fmt.Println(SumNumbers(bytes_array))

	fmt.Println("*********\n")
	Print(int_array)
	fmt.Println("*********")
	Print(float_array)
	fmt.Println("*********")
	Print(bytes_array)

	s1 := CustomSlice[int64]{32, 64, 96, 128}
	s2 := CustomSlice[float64]{32.0, 64.0, 96.1, 128.2}
	s3 := CustomSlice[int8]{8, 16, 24, 32}

	fmt.Println("*********")
	PrintCustomSliceElems(s1)
	PrintCustomSliceElems(s2)
	PrintCustomSliceElems(s3)

}
