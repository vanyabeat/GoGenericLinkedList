package LinkedList

import "fmt"

// listNode is a node of a doubly-linked list.
type listNode[T any] struct {
	// pointer to next node
	next *listNode[T]
	// pointer to prev node
	prev *listNode[T]
	// Value stored in node
	Value T
}

// LinkedList is a doubly-linked list.
// In cpp it is a template class.
// Empty list is a list with head and tail nodes.
/*
   prev          next           next
 ┌───────┐   ┌───────────┐    ┌───────┐
 │       │   │           │    │       │
 ▼      ┌┴───┴─┐        ┌▼────┴┐      ▼
nil  	│ head │        │ tail │     nil
        └────▲─┘        └┬─────┘
             │           │
             └───────────┘
                 prev
*/
type LinkedList[T any] struct {
	tail *listNode[T]
	head *listNode[T]
	Size uint64
}

// newListNode returns a new node with next and prev pointers set to nil and Value set to default value.
func newListNode[T any]() *listNode[T] {
	return &listNode[T]{next: nil, prev: nil, Value: *new(T)}
}

func newListNodeValueNextPrev[T any](p *listNode[T], v T, n *listNode[T]) *listNode[T] {
	return &listNode[T]{next: n, prev: p, Value: v}
}

func NewLinkedList[T any]() *LinkedList[T] {
	var list = new(LinkedList[T])
	list.Size = 0
	list.head = newListNode[T]()
	list.tail = newListNode[T]()
	list.head.next = list.tail
	list.tail.prev = list.head
	return list
}

func (l *LinkedList[T]) Empty() (result bool) {
	result = l.Size == 0
	return
}

func (l *LinkedList[T]) Len() int {
	return int(l.Size)
}

func (l *LinkedList[T]) PushBack(value T) {
	var node = newListNodeValueNextPrev[T](l.tail.prev, value, l.tail)
	l.tail.prev.next = node
	l.tail.prev = node
	l.Size += 1
	return
}

func (l *LinkedList[T]) PopBack() (result T) {
	if l.Empty() {
		panic("Empty list")
	}
	var node = l.tail.prev
	result = node.Value
	node.prev.next = l.tail
	l.tail.prev = node.prev
	l.Size -= 1
	return
}

func (l *LinkedList[T]) PopFront() (result T) {
	if l.Empty() {
		panic("Empty list")
	}
	var node = l.head.next
	result = node.Value
	node.next.prev = l.head
	l.head.next = node.next
	l.Size -= 1
	return
}

func (l *LinkedList[T]) PushFront(value T) {
	var node = newListNodeValueNextPrev[T](l.head, value, l.head.next)
	l.head.next.prev = node
	l.head.next = node
	l.Size += 1
	return
}

func (l *LinkedList[T]) Iterate() func() (T, bool) {
	var begin = l.head.next
	return func() (T, bool) {
		var value = begin.Value
		if begin == l.tail {
			return *new(T), false
		}
		begin = begin.next
		return value, true
	}
}

func (l *LinkedList[T]) At(index int) *T {
	if index < 0 || index >= int(l.Size) {
		panic("Index out of range")
	} else {
		var begin = l.head.next
		for i := 0; i < index; i++ {
			begin = begin.next
		}
		return &begin.Value
	}
}

func (l *LinkedList[T]) Index(index int) T {
	if index < 0 || index >= int(l.Size) {
		panic("Index out of range")
	} else {
		var begin = l.head.next
		for i := 0; i < index; i++ {
			begin = begin.next
		}
		return begin.Value
	}
}

func (l *LinkedList[T]) String() string {
	var result = "["
	var begin = l.head.next
	for begin != l.tail {
		result += fmt.Sprintf("%v", begin.Value)
		if begin.next != l.tail {
			result += " "
		}
		begin = begin.next
	}
	result += "]"
	return result
}

func (l *LinkedList[T]) Slice() []T {
	var result = make([]T, l.Size)
	var begin = l.head.next
	for i := 0; i < int(l.Size); i++ {
		result[i] = begin.Value
		begin = begin.next
	}
	return result
}
