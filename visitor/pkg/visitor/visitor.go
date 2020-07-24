package visitor

import (
	"errors"
	"fmt"

	"github.com/design-patterns/visitor/pkg/api/v1"
	"github.com/design-patterns/visitor/pkg/list"
	"github.com/design-patterns/visitor/pkg/queue"
	"github.com/design-patterns/visitor/pkg/stack"
)

// Visitor ...
type Visitor interface {
	JoinToList(l list.List) (err error)
	MoveToArray(s stack.Stack) (counter int, err error)
	InQueue(q queue.Queue) (ok bool, err error)
}

type array struct {
	slice []int
}

// JoinToList adds array of the elements to the List
func (a *array) JoinToList(l list.List) (err error) {
	if len(a.slice) == 0 {
		return errors.New(v1.ErrorMsgNothingToJoin)
	}
	counter := 0
	for _, item := range a.slice {
		l.Push(item)
		counter++
	}
	fmt.Println(v1.PushedToListMsgStart, counter, v1.PushedToListMsgEnd)
	return
}

// MoveToArray moves elements from the stack to the array until
// either the array runs out of space or the stack is empty.
// returns the number of moved elements
func (a *array) MoveToArray(s stack.Stack) (counter int, err error) {
	if len(a.slice) != 0 && a.slice[0] == 404 {
		// This is an artificial error, specially invented,
		// because it is necessary for this example
		err = errors.New(v1.ErrorMsg404)
		return
	}

	for len(a.slice) != cap(a.slice) {
		value, empty := s.Pop()
		if empty {
			break
		}
		a.slice = append(a.slice, value)
		counter++
	}
	return
}

// InQueue checks whether all elements from the array are in the queue
func (a *array) InQueue(q queue.Queue) (ok bool, err error) {
	if len(a.slice) == 0 {
		err = errors.New(v1.ErrorMsgNothingToSearch)
		return
	}

	for _, item := range a.slice {
		if ok = q.Search(item); !ok {
			return
		}
	}
	return
}

// NewArray visitor factory
func NewArray(slice []int) Visitor {
	return &array{
		slice,
	}
}
