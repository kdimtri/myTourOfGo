/*
Exercise: Errors
Copy your Sqrt function from the earlier exercise and modify it to return an error value.
Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.
Create a new type
type ErrNegativeSqrt float64
and make it an error by giving it a
func (e ErrNegativeSqrt) Error() string
method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".
Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?
Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.
*/
package main

import (
	"fmt"
)

// ErrNegativeSqrt error handeler in sqrt func
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if float64(e) < 0 {
		return fmt.Sprintf("\"cannot Sqrt negative number: %v\"\n", float64(e))
	}
	fmt.Printf("wrong call with nonNegotive value for ErrNegativeSqrt.Error(): %v\n", float64(e))
	panic(func() { return })
}

//Sqrt square root calculation
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	for y := 0; y < 10; y++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
