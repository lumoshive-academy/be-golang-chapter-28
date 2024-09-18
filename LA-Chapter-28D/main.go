// package main

// import "fmt"

// // Addable is an interface constraint that supports the + operator.
// type Addable interface {
//     int | int8 | int16 | int32 | int64 |
//     float32 | float64 |
//     string
// }

// // Sum sums the elements in a slice of Addable types.
// func Sum[T Addable](slice []T) T {
//     var sum T
//     for _, v := range slice {
//         sum += v
//     }
//     return sum
// }

// func main() {
//     intSlice := []int{1, 2, 3, 4, 5}
//     floatSlice := []float64{1.1, 2.2, 3.3}
//     stringSlice := []string{"Hello", " ", "world"}

//     fmt.Println("Sum of intSlice:", Sum(intSlice))        // Output: Sum of intSlice: 15
//     fmt.Println("Sum of floatSlice:", Sum(floatSlice))    // Output: Sum of floatSlice: 6.6
//     fmt.Println("Concatenation of stringSlice:", Sum(stringSlice)) // Output: Concatenation of stringSlice: Hello world
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// // Comparable is an interface that represents types that can be compared.
// type Comparable interface {
// 	~int | ~float64 | ~string
// }

// // Sort sorts a slice of Comparable elements in ascending order.
// func Sort[T Comparable](slice []T) []T {
// 	sort.Slice(slice, func(i, j int) bool {
// 		return slice[i] < slice[j]
// 	})
// 	return slice
// }

// func main() {
// 	intSlice := []int{5, 3, 9, 1, 4}
// 	floatSlice := []float64{3.2, 1.5, 4.8, 2.9}
// 	stringSlice := []string{"banana", "apple", "cherry"}

// 	// Sorting slices using Sort function
// 	fmt.Println("Sorted int slice:", Sort(intSlice))       // Output: Sorted int slice: [1 3 4 5 9]
// 	fmt.Println("Sorted float slice:", Sort(floatSlice))   // Output: Sorted float slice: [1.5 2.9 3.2 4.8]
// 	fmt.Println("Sorted string slice:", Sort(stringSlice)) // Output: Sorted string slice: [apple banana cherry]
// }

package main

import (
	"fmt"
)

// Define a generic function that accepts any type
func PrintValue[T any](value T) {
	fmt.Println(value)
}

func main() {
	// Call the generic function with different types of arguments
	PrintValue(42)           // T inferred as int
	PrintValue("Hello, Go!") // T inferred as string
	PrintValue(3.14)         // T inferred as float64
	PrintValue(true)         // T inferred as bool
}
