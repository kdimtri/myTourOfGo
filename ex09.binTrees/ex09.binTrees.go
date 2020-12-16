/*
Exercise: Equivalent Binary Trees
1. Implement the Walk function.
2. Test the Walk function.
	- The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.
	- Create a new channel ch and kick off the walker:
		go Walk(tree.New(1), ch)
	- Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.
3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.
4. Test the Same function.
	- Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.
*/
package main

import (
	"golang.org/x/tour/tree"

	"fmt"
)

//Walk walks the tree t sending all values from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
//func Same(t1, t2 *tree.Tree) bool

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		v, ok := <-ch
		fmt.Printf("%v) t(1)val: %v and ch is ok? %v\n", i, v, ok)
		if !ok {
			close(ch)
			break
		}
	}
}
