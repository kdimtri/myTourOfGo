/*
Exercise: Readers
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/
package main

import (
	//"io"

	"golang.org/x/tour/reader"
)

//MyReader fils any byte slice with ASCII character 'A'
type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	for i := range b {
		b[i] = byte('A')
	}
	return len(b), nil //shouldn't i return here an error="io.EOF"?
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
