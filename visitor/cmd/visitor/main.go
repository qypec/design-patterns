package main

import (
	"fmt"

	"github.com/design-patterns/visitor/pkg/list"
	"github.com/design-patterns/visitor/pkg/queue"
	"github.com/design-patterns/visitor/pkg/stack"
	"github.com/design-patterns/visitor/pkg/visitor"
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
	arrL.JoinToList(l)
	fmt.Println(l)

	arrS := visitor.NewArray([]int{100, 200, 300, 400, 500})
	arrS.JoinToStack(s)
	fmt.Println(s)

	arrQ := visitor.NewArray([]int{1000, 2000, 3000, 4000, 5000})
	arrQ.JoinToQueue(q)
	fmt.Println(q)
}
