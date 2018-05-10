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
	l1 := list.NewLinkedList()
	l2 := list.NewLinkedList()
	GetParam(l1)
	GetParam(l2)
}

// GetParam 获取计算参数
func GetParam(l *list.LinkedList) {
	var count int
	fmt.Scanf("%d", &count)
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
