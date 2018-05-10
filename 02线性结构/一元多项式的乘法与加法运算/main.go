package main

import (
	"fmt"
)

type item struct {
	coefficient int
	index       int
}
type pitem struct {
	item
	next *item
}
type polynomial struct {
	Head  *item
	count int
}

func (list *polynomial) Append(){

}

func main() {
	var count int
	fmt.Scanf("%d", &count)
	var p1 polynomial
	p1.count = count
	for i := 0; i < count; i++ {
		var it item
		var coefficient int
		var index int
		fmt.Scanf("%d", coefficient)
		fmt.Scanf("%d", index)
		it.index = index
		it.coefficient = coefficient
		p1.next = loc
		loc = &p1
	}
}
