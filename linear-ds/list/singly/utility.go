package singly

// ----------------- Utility methods -----------------

func (l *Singly[T]) Reverse() {

	if l.Head == nil {
		return
	}
	var prev *Node[T]
	curr := l.Head
	l.Tail = l.Head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}

// func (l *Singly[T]) ReverseCopy() Singly[T] {

// }

// func (l Singly[T]) Clone() *Singly[T] {

// }
