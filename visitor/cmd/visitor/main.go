package main

import (
	"fmt"

	"github.com/design-patterns/visitor/pkg/list"
	"github.com/design-patterns/visitor/pkg/queue"
	"github.com/design-patterns/visitor/pkg/stack"
	"github.com/design-patterns/visitor/pkg/visitor"
)

const (
	moveToArrayMsg = " items moved from stack to array\n"
	allInQueueMsg  = "All elements from the array are in the queue\n"
)

func main() {
	l := list.NewList()
	l.Push(1)
	l.Push(2)
	l.Push(3)

	s := stack.NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	q := queue.NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	arrL := visitor.NewArray([]int{10, 20, 30, 40, 50})
	err := l.Accept(arrL)
	if err != nil {
		panic(err)
	}
	fmt.Println(l)

	arrS := visitor.NewArray(make([]int, 0, 10))
	counter, err := s.Accept(arrS)
	if err != nil {
		panic(err)
	}
	fmt.Println(counter, moveToArrayMsg)

	arrQ := visitor.NewArray([]int{3, 1, 2})
	ok, err := q.Accept(arrQ)
	if err != nil {
		panic(err)
	}
	if ok {
		fmt.Println(allInQueueMsg)
	}
}
