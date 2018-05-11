package list

import (
	"testing"
)

func CheckedAppend(t *testing.T) {
	t.Log("Start Testing")
	//	l := Append()
}

func TestList(t *testing.T) {
	t.Log("Start Testing/n")
	l := NewLinkedList()
	n1 := NewINode(1, nil)
	n2 := NewINode(2, nil)
	n3 := NewINode(3, nil)
	l.Append(n1)
	l.Append(n2)
	l.Append(n3)
	l.Print()
	l.Delete(1)
	l.Print()
}
