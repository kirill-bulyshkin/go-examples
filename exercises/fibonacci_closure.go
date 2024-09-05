package exercises

import "fmt"

// Exercise: Fibonacci closure

// Let's have some fun with functions.
// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	previousValue := 0
	currentValue := 1

	return func() int {
		res := previousValue
		previousValue = currentValue
		currentValue = res + currentValue
		return res
	}
}

func FibonacciTest() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
