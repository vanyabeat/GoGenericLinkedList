package LinkedList

import (
	"testing"
)

func TestLinkedList_Empty(t *testing.T) {
	list := NewLinkedList[int]()
	switch {
	case list.Size != 0:
		t.Fatalf("Size is not 0")
	case list.head.next != list.tail:
		t.Fatalf("head.next is not tail")
	case list.tail.prev != list.head:
		t.Fatalf("tail.prev is not head")
	case !list.Empty():
		t.Fatalf("Empty() is true")
	}
}

func TestLinkedList_PushBack(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushBack(1)
	switch {
	case list.Size != 1:
		t.Fatalf("Size is not 1")
	case list.head.next != list.tail.prev:
		t.Fatalf("head.next is not tail.prev")
	case list.head.next.Value != 1:
		t.Fatalf("head.next.Value is not 1")
	case list.tail.prev.Value != 1:
		t.Fatalf("tail.prev.Value is not 1")
	}
}
