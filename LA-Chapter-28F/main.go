// package main

// import "fmt"

// // GenericInterface is a generic interface that defines a method Process.
// type GenericInterface[T any] interface {
// 	Process(input T) T
// }

// // GenericStruct is a generic struct that implements GenericInterface.
// type GenericStruct[T any] struct {
// 	Value T
// }

// // Process is a method that implements the GenericInterface, simply returning the input value.
// func (g GenericStruct[T]) Process(input T) T {
// 	// This simple implementation just returns the input value
// 	return input
// }
// func main() {
// 	// Using GenericStruct with int type
// 	intStruct := GenericStruct[int]{Value: 10}
// 	resultInt := intStruct.Process(20)
// 	fmt.Println("Processed int:", resultInt) // Output: Processed int: 20

// 	// Using GenericStruct with string type
// 	stringStruct := GenericStruct[string]{Value: "hello"}
// 	resultString := stringStruct.Process("world")
// 	fmt.Println("Processed string:", resultString) // Output: Processed string: world
// }

package main

import "fmt"

// Add sums two values of a numeric type.
func Add[T int | float64](a, b T) T {
	return a + b
}

func main() {
	sumInt := Add(5, 10)      // Works with int
	sumFloat := Add(2.5, 3.5) // Works with float64

	fmt.Printf("Sum (int): %d\n", sumInt)       // Output: Sum (int): 15
	fmt.Printf("Sum (float64): %f\n", sumFloat) // Output: Sum (float64): 6.000000
}
