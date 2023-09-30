# GoLinkedList
Generic LinkedList implementation for go in C++ style
Bidirectional linked list with push/pop front/back operations

[![Go Report Card](https://goreportcard.com/badge/github.com/vanyabeat/GoLinkedList)](https://goreportcard.com/report/github.com/vanyabeat/GoLinkedList)

[![Tests](https://github.com/vanyabeat/GoLinkedList/actions/workflows/Tests.yml/badge.svg)](https://github.com/vanyabeat/GoLinkedList/actions/workflows/Tests.yml)
[![codecov](https://codecov.io/gh/vanyabeat/GoLinkedList/graph/badge.svg?token=6MXG5KLBGF)](https://codecov.io/gh/vanyabeat/GoLinkedList)

## Usage

```go   
package main

import (
	"fmt"
	"github.com/vanyabeat/GoLinkedList/LinkedList"
)

func main() {
	var list = LinkedList.NewLinkedList[int]()
	list.PushBack(1)
	list.PushBack(2)
	list.PushFront(666)
	fmt.Println(list)
}
```
    
## Output

```bash
[666 1 2]
```

## PushBack

```go
var list = LinkedList.NewLinkedList[int]()
list.PushBack(1)
list.PushBack(2)
fmt.Println(list) // [1 2]
```

## PushFront

```go
var list = LinkedList.NewLinkedList[int]()
list.PushFront(1)
list.PushFront(2)
fmt.Println(list) // [2 1]
```

## PopBack
```go
var list = LinkedList.NewLinkedList[int]()
list.PushBack(1)
list.PushBack(2)
list.PopBack()
fmt.Println(list) // [1]
```

## PopFront
```go
var list = LinkedList.NewLinkedList[int]()
list.PushBack(1)
list.PushBack(2)
fmt.Println(list.PopFront()) // 1
fmt.Println(list) // [2]
```

## Slice
```go
var list = LinkedList.NewLinkedList[int]()
list.PushBack(1)
list.PushBack(2)
list.PushBack(3)
list.PushBack(4)
var slc = list.Slice()
```