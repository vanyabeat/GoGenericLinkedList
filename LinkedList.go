package LinkedList

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

type IteratorLinkedList[T any] struct {
	current *listNode[T]
	end     *listNode[T]
}

func newListNode[T any]() *listNode[T] {
	return &listNode[T]{next: nil, prev: nil, Value: *new(T)}
}

func newListNodeValueNextPrev[T any](p *listNode[T], v T, n *listNode[T]) *listNode[T] {
	return &listNode[T]{next: n, prev: p, Value: v}
}

func (reciever *LinkedList[T]) Empty() (result bool) {
	result = reciever.Size == 0
	return
}

func (reciever *LinkedList[T]) Len() int {
	return int(reciever.Size)
}

func (reciever *LinkedList[T]) PushBack(value T) {
	var node = newListNodeValueNextPrev[T](reciever.tail.prev, value, reciever.tail)
	reciever.tail.prev.next = node
	reciever.tail.prev = node
	reciever.Size += 1
	return
}

func (reciever *LinkedList[T]) PopBack() (result T) {
	if reciever.Empty() {
		panic("Empty list")
	}
	var node = reciever.tail.prev
	result = node.Value
	node.prev.next = reciever.tail
	reciever.tail.prev = node.prev
	reciever.Size -= 1
	return
}

func (reciever *LinkedList[T]) PopFront() (result T) {
	if reciever.Empty() {
		panic("Empty list")
	}
	var node = reciever.head.next
	result = node.Value
	node.next.prev = reciever.head
	reciever.head.next = node.next
	reciever.Size -= 1
	return
}

func (reciever *LinkedList[T]) PushFront(value T) {
	var node = newListNodeValueNextPrev[T](reciever.head, value, reciever.head.next)
	reciever.head.next.prev = node
	reciever.head.next = node
	reciever.Size += 1
	return
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
