/*
Exercise: Readers
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/
package main

import "golang.org/x/tour/reader"

//MyReader fils any byte slice with ASCII character 'A'
type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = byte('A')
	}
	return len(b), nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
