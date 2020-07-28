package stack

import (
	"strconv"

	"github.com/design-patterns/visitor/pkg/api/v1"
)

type visitor interface {
	MoveToArray(s Stack) (counter int, err error)
}

// Stack interface describes the behavior of a stack
type Stack interface {
	Push(value int)
	Pop() (value int, empty bool)
	Len() (len int)
	String() (str string)
	Accept(v visitor) (counter int, err error)
}

type stack struct {
	// The head element of the stack is the last element added
	head *elem
	len  int
}

// Accept accepts the visitor
func (s *stack) Accept(v visitor) (counter int, err error) {
	return v.MoveToArray(s)
}

// Push creates a new stack element with a value and pushes it to the end of the stack
func (s *stack) Push(value int) {
	s.push(newElem(value))
}

// Pop deletes the head element of the stack and returns its value
func (s *stack) Pop() (value int, empty bool) {
	return s.pop()
}

// String method for the Stringer interface
func (s *stack) String() (str string) {
	str = v1.StackLenMsgStart + strconv.Itoa(s.len) + v1.StackLenMsgEnd
	for tmp := s.head; tmp != nil; tmp = tmp.prev {
		str += strconv.Itoa(tmp.value) + v1.StackElementSeparator
	}
	if s.len == 0 && s.head == nil {
		str += v1.NilString
	}
	str += v1.LineBreak
	return
}

// Len returns the number of elements
func (s *stack) Len() (len int) {
	return s.len
}

func (s *stack) push(new elem) {
	if s.len != 0 {
		new.prev = s.head
	}
	s.head = &new
	s.len++
}

func (s *stack) pop() (value int, empty bool) {
	if s.len == 0 {
		empty = true
		return
	}
	value = s.head.value
	s.head = s.head.prev
	s.len--
	return
}

// NewStack creates an instance of the Stack
func NewStack() Stack {
	return &stack{}
}
