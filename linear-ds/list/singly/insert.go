package singly


// ----------------- Insertion methods -----------------

func (l *Singly[T]) InsertAt(index int, elem T) {
	
}

func (l *Singly[T]) PushFront() {
	
}

func (l *Singly[T]) InsertAfter(elem T, val T) error {
	temp := l.Head
	
	// Traverse to the first occurence of the value
	for temp != nil && temp.Val != elem {
		temp = temp.Next
	}

	// if temp is nil the element doesn't exist so returning error
	if temp == nil {
		return ElementNotFound
	}

	// Now we have found the element
	newNode := &Node[T]{Val: val, Next: temp.Next}
	temp.Next = newNode
	l.Len++

	// if element is the last element then moving tail
	if temp == l.Tail {
		l.Tail = newNode
	}

	return nil
}

func (l *Singly[T]) Append(elem T) {
	// Inserting element next to the tail
	newNode := Node[T]{Val: elem, Next: nil}

	// Appending at the end 
	l.Tail.Next = &newNode

	// Taking tail to last element
	l.Tail = &newNode
}
