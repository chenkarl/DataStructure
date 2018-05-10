package list

import (
	"reflect"
)

// ElementType 代表所有类型
type ElementType interface{}

// INode 链表的结点
type INode struct {
	X    ElementType
	next *INode
}

// LinkedList 链表，其中包括一个头结点地址
// 头结点的X值代表该链表长度
type LinkedList struct {
	Head *INode
}

// NewINode 新建一个结点
func NewINode(x ElementType, next *INode) *INode {
	return &INode{x, next}
}

// NewLinkedList 新建一个链表
func NewLinkedList() *LinkedList {
	head := &INode{0, nil}
	return &LinkedList{head}
}

// IsEmpty 判断一个链表是否为空
func (list *LinkedList) IsEmpty() bool {
	return list.Head.next == nil
}

// Length 获取一个链表长度
func (list *LinkedList) Length() int {
	return int(reflect.ValueOf(list.Head.X).Int())
}

// Append 在链表末尾增加一个结点
func (list *LinkedList) Append(node *INode) {
	current := list.Head
	for {
		if current.next == nil {
			break
		}
		current = current.next
	}
	current.next = node
}

// Prepend 在链表头部增加一个结点
func (list *LinkedList) Prepend(node *INode) {
	node.next = list.Head
	list.Head = node
}

// Find 查找链表中的某一个元素
func (list *LinkedList) Find(x ElementType) (*INode, bool) {
	current := list.Head
	for {
		if current != nil {
			if current.X == x {
				return current, true
			}
			current = current.next
		} else {
			break
		}
	}
	return nil, false
}

// Delete 删除链表中的某一个结点
func (list *LinkedList) Delete(x ElementType) error {
	current := list.Head
	for {
		if current.next != nil {
			if current.X == x {
				current.next = current.next.next
				return nil
			}
			current = current.next
		} else {
			break
		}
	}
	return nil
}
