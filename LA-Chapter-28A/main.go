package main

import "fmt"

// SumInts sums the elements in a slice of integers.
func SumInts(ints []int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}

// SumStrings concatenates the elements in a slice of strings.
func SumStrings(strings []string) string {
	result := ""
	for _, v := range strings {
		result += v
	}
	return result
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"hello", " ", "world"}

	fmt.Println("Sum of integers:", SumInts(intSlice))                // Output: Sum of integers: 15
	fmt.Println("Concatenation of strings:", SumStrings(stringSlice)) // Output: Concatenation of strings: hello world
}
