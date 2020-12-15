/*
Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.
You might find strings.Fields helpful.
*/
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount counts words in a string
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
