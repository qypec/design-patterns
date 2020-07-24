package queue

import (
	"strconv"

	"github.com/design-patterns/visitor/pkg/api/v1"
)

type visitor interface {
	InQueue(q Queue) (ok bool, err error)
}

// Queue interface describes the behavior of a queue
type Queue interface {
	Push(value int)
	Pop() (value int)
	Len() (len int)
	String() (str string)
	Search(toFind int) (ok bool)
	Accept(v visitor) (ok bool, err error)
}

type queue struct {
	head *elem
	len  int
}

// Accept accepts the visitor
func (q *queue) Accept(v visitor) (ok bool, err error) {
	return v.InQueue(q)
}

// Push creates a new queue element with a value and pushes it to the end of the queue
func (q *queue) Push(value int) {
	q.push(newElem(value))
}

// Pop deletes the head element of the queue and returns its value
func (q *queue) Pop() (value int) {
	return q.pop()
}

// String method for the Stringer interface
func (q *queue) String() (str string) {
	str = v1.QueueLenMsgStart + strconv.Itoa(q.len) + v1.QueueLenMsgEnd
	for tmp := q.head; tmp != nil; tmp = tmp.next {
		str += strconv.Itoa(tmp.value) + v1.QueueElementSeparator
	}
	if q.len == 0 && q.head == nil {
		str += v1.NilString
	}
	str += v1.LineBreak
	return
}

// Len returns the number of elements
func (q *queue) Len() (len int) {
	return q.len
}

// Search searches items value in the queue
func (q *queue) Search(toFind int) (ok bool) {
	for tmp := q.head; tmp != nil; tmp = tmp.next {
		if tmp.value == toFind {
			return true
		}
	}
	return
}

func (q *queue) push(new *elem) {
	tmp := &q.head
	for *tmp != nil {
		tmp = &(*tmp).next
	}
	*tmp = new
	q.len++
}

func (q *queue) pop() (value int) {
	if q.len != 0 {
		value = q.head.value
		q.head = q.head.next
		q.len--
	}
	return
}

// NewQueue creates an instance of the Queue
func NewQueue() Queue {
	return &queue{
		nil,
		0,
	}
}
