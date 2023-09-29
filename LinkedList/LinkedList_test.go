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

func TestLinkedList_PushFront(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushFront(1)
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

func TestLinkedList_PushFrontTwoItems(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushFront(1)
	list.PushFront(666)
	switch {
	case list.Len() != 2:
		t.Fatalf("Len() is not 2")
	}
}

func TestLinkedList_Iterate(t *testing.T) {
	list := NewLinkedList[int]()
	var size int = 100

	for i := 0; i < size; i++ {
		list.PushFront(i)
	}

	var reverse []int
	var actual []int

	var iter = list.Iterate()
	for i, ok := iter(); ok; i, ok = iter() {
		actual = append(actual, i)
	}
	for i := size - 1; i >= 0; i-- {
		reverse = append(reverse, i)
	}
	if len(actual) != len(reverse) {
		t.Fatalf("len(actual) != len(reverse)")
	}
	// Compare actual and reverse
	for i := 0; i < len(actual); i++ {
		if actual[i] != reverse[i] {
			t.Fatalf("actual[i] != reverse[i]")
		}
	}
}

func TestLinkedList_PopBack(t *testing.T) {
	list := NewLinkedList[int]()
	var size int = 1000
	for i := 0; i < size; i++ {
		list.PushFront(i)
	}
	for i := 0; i < size; i++ {
		if list.PopBack() != i {
			t.Fatalf("list.PopBack() != i")
		}
	}
}

func TestLinkedList_PopFront(t *testing.T) {
	list := NewLinkedList[int]()
	var size int = 1000
	for i := 0; i < size; i++ {
		list.PushFront(i)
	}
	for i := size - 1; i != 0; i-- {
		var value int = list.PopFront()
		if value != i {
			t.Fatalf("list.PopBack(%d) != %d", i, value)
		}
	}
}

func TestLinkedList_PopFrontEmpty(t *testing.T) {
	list := NewLinkedList[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic")
		}
	}()
	list.PopFront()
}

func TestLinkedList_PopBackEmpty(t *testing.T) {
	list := NewLinkedList[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic")
		}
	}()
	list.PopBack()
}

func TestLinkedList_At(t *testing.T) {
	list := NewLinkedList[int]()
	list.PushFront(1)
	*list.At(0) = 666
	if list.head.next.Value != 666 {
		t.Fatalf("list.head.next.Value != 1")
	}
}

func TestLinkedList_AtPanic(t *testing.T) {
	list := NewLinkedList[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic")
		}
	}()
	list.At(0)
}
