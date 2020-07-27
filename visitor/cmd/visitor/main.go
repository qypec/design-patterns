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

	errorVisitingList = "Error while visiting the List: "
)

func main() {
	l := list.NewList()
	s := stack.NewStack()
	q := queue.NewQueue()

	for i := 0; i < 3; i++ {
		l.Push(i + 1)
		s.Push(i + 1)
		q.Push(i + 1)
	}

	arrL := visitor.NewArray([]int{10, 20, 30, 40, 50})
	err := l.Accept(arrL)
	if err != nil {
		fmt.Println(errorVisitingList, err)
	}
	fmt.Println(l)

	arrS := visitor.NewArray(make([]int, 0, 10))
	counter, err := s.Accept(arrS)
	if err != nil {
		fmt.Println(errorVisitingList, err)
	}
	fmt.Println(counter, moveToArrayMsg)

	arrQ := visitor.NewArray([]int{3, 1, 2})
	ok, err := q.Accept(arrQ)
	if err != nil {
		fmt.Println(errorVisitingList, err)
	}
	if ok {
		fmt.Println(allInQueueMsg)
	}
}
