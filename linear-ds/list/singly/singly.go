package singly

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

type Singly[T any] struct {
	Head *Node[T]
	Len int
}

// Factory Function
func New[T any]() *Singly[T] {
	return &Singly[T]{Head: nil, Len: 0}
}


