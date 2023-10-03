package exercises

import (
	"fmt"
	"math"
)

//Exercise: Loops and Functions

// As a way to play with functions and loops, let's implement a square root function: given a number x, we want to find the number z for which z² is most nearly x.

// Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:
// z -= (z*z - x) / (2*z)

func Sqrt(x float64) float64 {

	var z, size float64 = x / 2, 10000

	for i := 0; i <= 10; i++ {
		fmt.Printf("Iteration_%d: ", i)
		// Comparing current rounded value of z with the rounded value after one more attemting of sqrtFormula()
		if z != RoundToSize(sqrtFormula(x, z), size) {

			z = RoundToSize(sqrtFormula(x, z), size)
			fmt.Printf("%.15f\n", z)

		}
	}

	fmt.Printf("\n\nSquare root of %f is %f", x, z)
	return z
}

func sqrtFormula(x, z float64) float64 {
	z -= (z*z - x) / (2 * z)
	return z
}

func RoundToSize(x float64, size float64) float64 {
	roundedFloat := math.Round(x*size) / size
	return roundedFloat
}
