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

func (l *Singly[T]) ForEach(f func(*Node[T])) {
	temp := l.Head

	for temp != nil {
		f(temp)
		temp = temp.Next
	}
}

// func (l Singly[T])ToSlice() []T {

// }
