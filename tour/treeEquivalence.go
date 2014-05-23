package main

import (
	"fmt"

	"code.google.com/p/go-tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	if nil != t.Left {
		walk(t.Left, ch)
	}

	ch <- t.Value
	if nil != t.Right {
		walk(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}

		if !(ok1 && ok2) {
			break
		}
	}

	return true
}

func main() {
	size := 5
	ch := make(chan int)
	go Walk(tree.New(size), ch)
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
