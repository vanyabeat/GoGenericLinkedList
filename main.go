package main

import (
	"fmt"
	"github.com/vanyabeat/GoLinkedList/LinkedList"
)

func main() {
	var list = LinkedList.NewLinkedList[int]()
	list.PushBack(1)
	list.PushBack(2)
	fmt.Println(list)
}
