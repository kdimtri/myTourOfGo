/*
Exercise: Loops and Functions
As a way to play with functions and loops, let's implement a square root function: given a number x, we want to find the number z for which z² is most nearly x.
Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess: ....
*/
package main

import (
	"fmt"
)

//Sqrt square root calculation
func Sqrt(x float64) float64 {
	z := 1.0
	for y := 0; y < 10; y++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(4))
}
