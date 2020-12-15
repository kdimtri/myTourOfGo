/*
Exercise: Slices
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
(You need to use a loop to allocate each []uint8 inside the [][]uint8.)
(Use uint8(intValue) to convert between types.)
*/
package main

import "golang.org/x/tour/pic"

// Pic function that generates image array
func Pic(dx, dy int) [][]uint8 {
	r := make([][]uint8, dy)
	for i := range r {
		r[i] = make([]uint8, dx)
		for j := range r[i] {
			r[i][j] = uint8((i ^ j) * (i * j) / 4)
		}
	}
	return r
}

func main() {
	pic.Show(Pic)
}
