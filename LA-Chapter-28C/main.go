// package main

// import "fmt"

// // FindIndex searches for an element in a slice and returns its index.
// // If the element is not found, it returns -1.
// func FindIndex[T comparable](slice []T, value T) int {
// 	for i, v := range slice {
// 		if v == value {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func main() {
// 	intSlice := []int{1, 2, 3, 4, 5}
// 	stringSlice := []string{"apple", "banana", "cherry"}

// 	// Searching for an integer in a slice of integers
// 	index := FindIndex(intSlice, 3)
// 	fmt.Printf("Index of 3 in intSlice: %d\n", index) // Output: Index of 3 in intSlice: 2

// 	// Searching for a string in a slice of strings
// 	index = FindIndex(stringSlice, "banana")
// 	fmt.Printf("Index of 'banana' in stringSlice: %d\n", index) // Output: Index of 'banana' in stringSlice: 1

// 	// Searching for a string not present in the slice
// 	index = FindIndex(stringSlice, "orange")
// 	fmt.Printf("Index of 'orange' in stringSlice: %d\n", index) // Output: Index of 'orange' in stringSlice: -1
// }

package main

import (
	"fmt"
)

// Define an interface that will be used as a type constraint
type Stringer interface {
	String() string
}

// Define a struct that implements the Stringer interface
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

// Define a generic function that works with any type that implements the Stringer interface
func PrintString[T Stringer](value T) {
	fmt.Println(value.String())
}

func main() {
	p := Person{Name: "Lumoshive", Age: 20}
	PrintString(p)
}
