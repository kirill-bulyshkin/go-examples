// go run examples.go - to run Go program
// go build examples.go - to build binary Go file

package main

import (
	"fmt"
)

func main() {
	appendElementsToSlice()
}

// DECLARING VARIABLES
func varWithInitialValue() {
	var student1 string = "John" // type is string
	var student2 = "Jane"        // type is inferred
	x := 2                       // type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}

func varWOInitialValue() {
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func varAssignmentAfterDeclaration() {
	var student1 string
	student1 = "John"

	fmt.Println(student1)
}

func multipleVariableDeclaration() {
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func multipleVariableDeclarationWOType() {
	var a, b = 6, "Hello"
	c, d := 7, "World!"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func multipleVariableDeclarationInBlock() {
	var (
		a int
		b int    = 1
		c string = "hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// OUTPUT FUNCTIONS
// - Print() - prints its arguments with their default format
// - Println() - a whitespace is added between the arguments, and a newline is added at the end
// - Printf() - with arguments based on the given formatting verb

func print() {
	var i, j string = "Hello", "World"

	fmt.Print(i, "\n")
	fmt.Print(j, "\n")
}

func println() {
	var i, j string = "Hello", "World"

	fmt.Println(i, j)
}

func printf() {
	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}

//DATA TYPES
//bool
//string
//Numeric:

// - int
// - int8
// - int16
// - int32
// - int64

// - uint
// - uint8
// - uint16
// - uint32
// - uint64

// - float32
// - float64

// ARRAYS
// used to store multiple values of the same type
// you cannot change the length of an array
func arrayWithDefinedLength() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}

func arrayWithInferredLength() {
	var arr1 = [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}

func accessElementsOfArray() {
	prices := [3]int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])
}

func changeElementsOfArrayValue() {
	prices := [3]int{10, 20, 30}

	prices[2] = 50
	fmt.Println(prices)
}

func arrayInitialization() {
	arr1 := [5]int{}              // not initialized
	arr2 := [5]int{1, 2}          // partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} // fully initialized

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

func arrayInitializeOnlySpecificElements() {
	arr1 := [5]int{1: 10, 2: 40} // 1 and 2 are indexes of the elements

	fmt.Println(arr1)
}

func findLengthOfArray() {
	arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr2 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
}

// SLICES
// similar to arrays, but are more flexible
// you can change the length of a slice

// How to create slice:
// - Using the []datatype{values} format
// - Create a slice from an array
// - Using the make() function

func createSliceBasicWay() {
	myslice1 := []int{}
	fmt.Println(len(myslice1)) // len() function - returns the length of the slice (the number of elements in the slice)
	fmt.Println(cap(myslice1)) // cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
}

// var myarray = [length]datatype{values} // An array
// myslice := myarray[start:end] // A slice made from the array

// The slice starts from the second element of the array which has value 12.
// The slice can grow to the end of the array.
// This means that the capacity of the slice is 4.
func createSliceFromArray() {
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
}

// slice_name := make([]type, length, capacity)
func createSliceUsingMake() {
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// with omitted capacity
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))
}

// You can append elements to the end of a slice using the append()function:
// slice_name = append(slice_name, element1, element2, ...)
// it will cnahge the slice length
func appendElementsToSlice() {
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
}

// slice3 = append(slice1, slice2...)
func appendOneSliceToAnotherSlice() {
	myslice1 := []int{1, 2, 3}
	myslice2 := []int{4, 5, 6}
	myslice3 := append(myslice1, myslice2...)

	fmt.Printf("myslice3=%v\n", myslice3)
	fmt.Printf("length=%d\n", len(myslice3))
	fmt.Printf("capacity=%d\n", cap(myslice3))
}

// To change legth of slice need to add or remove some items
func сhangeLengthOfSlice() {
	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice1 := arr1[1:5]                 // Slice array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = arr1[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
}

// If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.
// copy(dest, src)
func copyItemsFromOneSliceToAnother() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}

//OPERATORS

//Arithmetic Operators
//   +	Addition
//   -	Subtraction
//   *	Multiplication
//   /	Division
//   % Modulus
//   ++ Increment
//   -- Decrement

// Assignment Operators
//	=
//	+=
// 	-=
//	...

//Comparison Operators
// ==	Equal to
// !=	Not equal
// >	Greater than
// <	Less than
// >=	Greater than or equal to
// <=	Less than or equal to

//Logical Operators
// && 	Logical and	Returns true if both statements are true
// || 	Logical or	Returns true if one of the statements is true
// !	Logical not Reverse the result, returns false if the result is true

//Bitwise Operators
// & 	AND	Sets each bit to 1 if both bits are 1	x & y
// |	OR	Sets each bit to 1 if one of two bits is 1	x | y
//  ^	XOR	Sets each bit to 1 if only one of two bits is 1	x ^ b
// <<	Zero fill left shift	Shift left by pushing zeros in from the right	x << 2
// >>	Signed right shift	Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off	x >> 2

// CONDITIONS
func elseIfStatement() {
	time := 22
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}
}

// SWITCH
// Go ONLY runs the matched case so it does not need a break statement.
func multiCaseSwitch() {
	day := 5

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}

// LOOPS
// Continue statement is used to skip one or more iterations in the loop.
func loopWithContinueStatement() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}

// The break statement is used to break/terminate the loop execution.
func loopWithBreakStatement() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}

// "Inner loop" will be executed one time for each iteration of the "outer loop"
func nestedLoops() {
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}
}

//The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value.
//TBU
