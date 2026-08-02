package singly

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
}

type Singly[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
	Len  int
}

// Factory Function
func New[T comparable]() *Singly[T] {
	return &Singly[T]{Head: nil, Len: 0}
}
