package main

import "fmt"

// Pair represents a pair of values of different types.
type Pair[T, U any] struct {
	First  T
	Second U
}

// NewPair creates a new Pair from two values of potentially different types.
func NewPair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{First: first, Second: second}
}

func main() {
	// Creating a pair of int and string
	p1 := NewPair(1, "one")
	fmt.Printf("Pair 1: (%v, %v)\n", p1.First, p1.Second)

	// Creating a pair of float64 and bool
	p2 := NewPair(3.14, true)
	fmt.Printf("Pair 2: (%v, %v)\n", p2.First, p2.Second)
}
