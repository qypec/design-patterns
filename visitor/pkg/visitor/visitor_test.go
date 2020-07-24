package visitor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/design-patterns/visitor/pkg/list"
	"github.com/design-patterns/visitor/pkg/queue"
	"github.com/design-patterns/visitor/pkg/stack"
)

const (
	methodNameJoinToList  = "JoinToList"
	methodNameJoinToStack = "JoinToStack"
	methodNameJoinToQueue = "JoinToQueue"

	expectedJoinToNonEmptyList  = "List (len = 5): 5 4 3 2 1 \n"
	expectedJointToEmptyList    = "List (len = 4): 5 4 3 2 \n"
	expectedJoinToNonEmptyStack = "Stack (len = 5): 5 4 3 2 1 \n"
	expectedJointToEmptyStack   = "Stack (len = 4): 5 4 3 2 \n"
	expectedJoinToNonEmptyQueue = "Queue (len = 5): 1 2 3 4 5 \n"
	expectedJointToEmptyQueue   = "Queue (len = 4): 2 3 4 5 \n"
)

func TestArray_JoinToNonEmptyList(t *testing.T) {
	testCase := []int{2, 3, 4, 5}

	t.Run(methodNameJoinToList, func(t *testing.T) {
		l := list.NewList()
		l.Push(1)
		array := NewArray(testCase)
		array.JoinToList(l)
		actual := l.String()
		assert.Equal(t, expectedJoinToNonEmptyList, actual)
	})
}

func TestArray_JoinToEmptyList(t *testing.T) {
	testCase := []int{2, 3, 4, 5}

	t.Run(methodNameJoinToList, func(t *testing.T) {
		l := list.NewList()
		array := NewArray(testCase)
		array.JoinToList(l)
		actual := l.String()
		assert.Equal(t, expectedJointToEmptyList, actual)
	})
}

func TestArray_JoinToNonEmptyStack(t *testing.T) {
	testCase := []int{2, 3, 4, 5}

	t.Run(methodNameJoinToStack, func(t *testing.T) {
		s := stack.NewStack()
		s.Push(1)
		array := NewArray(testCase)
		array.JoinToStack(s)
		actual := s.String()
		assert.Equal(t, expectedJoinToNonEmptyStack, actual)
	})
}

func TestArray_JoinToEmptyStack(t *testing.T) {
	testCase := []int{2, 3, 4, 5}

	t.Run(methodNameJoinToStack, func(t *testing.T) {
		s := stack.NewStack()
		array := NewArray(testCase)
		array.JoinToStack(s)
		actual := s.String()
		assert.Equal(t, expectedJointToEmptyStack, actual)
	})
}

func TestArray_JoinToNonEmptyQueue(t *testing.T) {
	testCase := []int{2, 3, 4, 5}

	t.Run(methodNameJoinToQueue, func(t *testing.T) {
		q := queue.NewQueue()
		q.Push(1)
		array := NewArray(testCase)
		array.JoinToQueue(q)
		actual := q.String()
		assert.Equal(t, expectedJoinToNonEmptyQueue, actual)
	})
}

func TestArray_JoinToEmptyQueue(t *testing.T) {
	testCase := []int{2, 3, 4, 5}

	t.Run(methodNameJoinToQueue, func(t *testing.T) {
		q := queue.NewQueue()
		array := NewArray(testCase)
		array.JoinToQueue(q)
		actual := q.String()
		assert.Equal(t, expectedJointToEmptyQueue, actual)
	})
}
