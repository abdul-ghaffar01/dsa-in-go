package singly


// ----------------- Insertion methods -----------------

func (l *Singly[T]) InsertAt(index int, elem T) {
	
}

func (l *Singly[T]) PushFront() {
	
}

func (l *Singly[T]) InsertAfter() {
	
}

func (l *Singly[T]) Append(elem T) {
	// Inserting element next to the tail
	newNode := Node[T]{Val: elem, Next: nil}

	// Appending at the end 
	l.Tail.Next = &newNode

	// Taking tail to last element
	l.Tail = &newNode
}
