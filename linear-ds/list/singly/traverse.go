package singly

import "fmt"

// ----------------- Traversal methods -----------------

func (l Singly[T]) Print() {
	for l.Head != nil {
		fmt.Printf("%v ", l.Head.Val)
		l.Head = l.Head.Next
	}
	fmt.Printf("\n")
}

// func (l *Singly[T]) ForEach(func(*Node[T])) {

// }

// func (l Singly[T])ToSlice() []T {

// }
