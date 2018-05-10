package main

import (
	"fmt"

	"github.com/chenkarl/DataStructure/scr/list"
)

type item struct {
	coefficient int
	index       int
}

func main() {
	var count int
	fmt.Scanf("%d", &count)
	l := list.NewLinkedList()
	for i := 0; i < count; i++ {
		var coefficient int
		var index int
		fmt.Scanf("%d", coefficient)
		fmt.Scanf("%d", index)
		var it item
		it.coefficient = coefficient
		it.index = index
		newNode := list.NewINode(it, nil)
		l.Append(newNode)
	}
}
