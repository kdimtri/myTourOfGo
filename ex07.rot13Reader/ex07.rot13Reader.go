/*
Exercise: rot13Reader
A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).
Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.
The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
*/
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func isAlfa(b byte) bool {
	return ((b >= 'A') && (b <= 'Z')) || ((b >= 'a') && (b <= 'z'))
}
func rot13(b byte) byte {
	if isAlfa(b) != true {
		return b
	}
	a0 := byte('a')
	if b < a0 {
		a0 = byte('A')
	}
	return (b-a0+13)%26 + a0
}

func (rr rot13Reader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	nb, err := rr.r.Read(b)
	for i := range b {
		b[i] = rot13(b[i])
	}
	return nb, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	s1 := &r //strings.NewReader()
	bs := make([]byte, 100)
	r1 := rot13Reader{
		r: r,
	}
	_, err := s1.Read(bs)
	if err == nil {
		fmt.Println(bs)
		fmt.Println(len(bs))
	}
	io.Copy(os.Stdout, &r)
	io.Copy(os.Stdout, &r1)
}
