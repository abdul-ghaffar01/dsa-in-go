package singly

// ----------------- Insertion methods -----------------

func (l *Singly[T]) InsertAt(index int, elem T) error {

	if index > l.Len {
		return OutOfRangeErr
	}
	// if index is 0
	if index == 0 {
		l.PushFront(elem)
		return nil
	}

	// if index is equal to len of list
	if index == l.Len {
		l.Append(elem)
		return nil
	}

	newNode := &Node[T]{Val: elem, Next: nil}
	// move to the correct index
	temp := l.Head
	for i := 0; i < index; i++ {
		temp = temp.Next
	}
	newNode.Next = temp.Next
	temp.Next = newNode
	l.Len++

	return nil
}

func (l *Singly[T]) PushFront(val T) {
	newNode := &Node[T]{Val: val, Next: l.Head}

	l.Head = newNode
	l.Len++

}

func (l *Singly[T]) InsertAfter(elem T, val T) error {
	temp := l.Head

	// Traverse to the first occurence of the value
	for temp != nil && temp.Val != elem {
		temp = temp.Next
	}

	// if temp is nil the element doesn't exist so returning error
	if temp == nil {
		return ElementNotFoundErr
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

	l.Len++
}
