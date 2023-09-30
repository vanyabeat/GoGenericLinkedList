package main

import (
	"fmt"
	"github.com/vanyabeat/GoLinkedList/LinkedList"
)

func main() {
	var list = LinkedList.NewLinkedList[interface{}]()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack("WTG")
	list.PushBack(true)
	list.PushBack(666)
	fmt.Println(list)
}
