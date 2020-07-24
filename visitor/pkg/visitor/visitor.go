package visitor

import (
	"fmt"

	"github.com/design-patterns/visitor/pkg/api/v1"
	"github.com/design-patterns/visitor/pkg/list"
	"github.com/design-patterns/visitor/pkg/queue"
	"github.com/design-patterns/visitor/pkg/stack"
)

// Visitor ...
type Visitor interface {
	JoinToList(l list.List)
	JoinToStack(s stack.Stack)
	JoinToQueue(q queue.Queue)
}

type array struct {
	slice []int
}

// JoinToList adds array of the elements to the List
func (a *array) JoinToList(l list.List) {
	counter := 0
	for _, item := range a.slice {
		l.Push(item)
		counter++
	}
	fmt.Println(v1.PushedToListMsgStart, counter, v1.PushedToListMsgEnd)
}

// JoinToStack adds array of the elements to the Stack
func (a *array) JoinToStack(s stack.Stack) {
	counter := 0
	for _, item := range a.slice {
		s.Push(item)
		counter++
	}
	fmt.Println(v1.PushedToStackMsgStart, counter, v1.PushedToStackMsgEnd)
}

// JoinToQueue adds array of the elements to the Queue
func (a *array) JoinToQueue(q queue.Queue) {
	counter := 0
	for _, item := range a.slice {
		q.Push(item)
		counter++
	}
	fmt.Println(v1.PushedToQueueMsgStart, counter, v1.PushedToQueueMsgEnd)
}

// NewArray factory
func NewArray(slice []int) Visitor {
	return &array{
		slice,
	}
}
