package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func WalkImpl(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		WalkImpl(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		WalkImpl(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkImpl(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1ch := make(chan int)
	t2ch := make(chan int)
	go Walk(t1, t1ch)
	go Walk(t2, t2ch)
	for {
		v1, ok1 := <-t1ch
		v2, ok2 := <-t2ch
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(3), ch)
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
