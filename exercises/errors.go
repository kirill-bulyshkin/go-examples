package exercises

import "fmt"

// Exercise: Errors

// Copy your Sqrt function from the earlier exercise and modify it to return an error value.
// Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

// Create a new type
// type ErrNegativeSqrt float64
// and make it an error by giving it a
// func (e ErrNegativeSqrt) Error() string
// method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".

// Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?
// Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	convertedValue := fmt.Sprint(float64(e))
	errStr := fmt.Sprintf("cannot Sqrt negative number: %s", convertedValue)
	return errStr
}

func SqrtWithErr(x float64) (float64, error) {
	if x < 0 {
		errStr := ErrNegativeSqrt(x).Error()
		fmt.Printf("\n error: %s", errStr)
		return 0, ErrNegativeSqrt(x)
	}

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
	return z, nil
}
