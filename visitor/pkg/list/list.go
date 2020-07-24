package list

import (
	"strconv"

	"github.com/design-patterns/visitor/pkg/api/v1"
)

type visitor interface {
	JoinToList(l List)
}

// List interface describes the behavior of a single-linked list
type List interface {
	Push(value int)
	Pop() (value int)
	Len() (len int)
	String() (str string)
}

type elem struct {
	value int
	next  *elem
}

type list struct {
	head *elem
	len  int
}

// Accept accepts the visitor
func (l *list) Accept(v visitor) {
	v.JoinToList(l)
}

// Push creates a new list element with a value and pushes it to the head of the list
func (l *list) Push(value int) {
	l.push(newElem(value))
}

// Pop deletes the head element of the list and returns its value
func (l *list) Pop() (value int) {
	return l.pop()
}

// String method for the Stringer interface
func (l *list) String() (str string) {
	str = v1.ListLenMsgStart + strconv.Itoa(l.len) + v1.ListLenMsgEnd
	for tmp := l.head; tmp != nil; tmp = tmp.next {
		str += strconv.Itoa(tmp.value) + v1.ListElementSeparator
	}
	if l.len == 0 && l.head == nil {
		str += v1.NilString
	}
	str += v1.LineBreak
	return
}

// Len returns the number of elements
func (l *list) Len() (len int) {
	return l.len
}

func (l *list) push(new *elem) {
	if l.len != 0 {
		new.next = l.head
	}
	l.head = new
	l.len++
}

func (l *list) pop() (value int) {
	if l.len != 0 {
		value = l.head.value
		l.head = l.head.next
		l.len--
	}
	return
}

func newElem(value int) *elem {
	return &elem{
		value: value,
		next:  nil,
	}
}

// NewList creates an instance of the List
func NewList() List {
	return &list{
		nil,
		0,
	}
}
