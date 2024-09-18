// package main

// import "fmt"

// // Box is a generic type that holds a value of any type.
// type Box[T any] struct {
// 	Value T
// }

// // IntBox is a specific type for Box holding an int.
// type IntBox Box[int]

// // StringBox is a specific type for Box holding a string.
// type StringBox Box[string]

// func main() {
// 	// Create an IntBox with an integer value
// 	intBox := IntBox{Value: 123}
// 	fmt.Printf("IntBox: %+v, Value: %d\n", intBox, intBox.Value)

// 	// Create a StringBox with a string value
// 	stringBox := StringBox{Value: "hello"}
// 	fmt.Printf("StringBox: %+v, Value: %s\n", stringBox, stringBox.Value)
// }

package main

import "fmt"

// Pair is a generic struct holding two values of potentially different types.
type Pair[T, U any] struct {
	First  T
	Second U
}

// SetValues sets the values of First and Second fields.
func (p *Pair[T, U]) SetValues(first T, second U) {
	p.First = first
	p.Second = second
}

func main() {
	// Using the generic Pair with int and string types
	pair := Pair[int, string]{First: 1, Second: "one"}
	fmt.Printf("Before SetValues: %+v\n", pair)

	// Set new values for the Pair
	pair.SetValues(2, "two")
	fmt.Printf("After SetValues: %+v\n", pair)
}
