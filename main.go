package main

import (
	"dsa-in-go/linear-ds/list/singly"
	"fmt"
)

func main() {
	list := singly.New[int]()
	list.Append(4)
	list.InsertAt(1, 9)
	list.Print()
	fmt.Println(list.Len)
}
