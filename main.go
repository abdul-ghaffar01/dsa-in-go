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
	
	//	for each loop 
	list.ForEach(func(n *singly.Node[int]) {
		fmt.Println(n.Val * n.Val)
	})
}
