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
	methodNameMoveToArray = "MoveToArray"
	methodNameInQueue     = "InQueue"

	expectedJoinToNonEmptyList = "List (len = 5): 5 4 3 2 1 \n"
	expectedJointToEmptyList   = "List (len = 4): 5 4 3 2 \n"
)

func TestArray_JoinToNonEmptyList(t *testing.T) {
	testCase := []int{2, 3, 4, 5}

	t.Run(methodNameJoinToList, func(t *testing.T) {
		l := list.NewList()
		l.Push(1)
		array := NewArray(testCase)
		err := array.JoinToList(l)
		actual := l.String()
		assert.Equal(t, expectedJoinToNonEmptyList, actual)
		assert.NoError(t, err)
	})
}

func TestArray_JoinToEmptyList(t *testing.T) {
	testCase := []int{2, 3, 4, 5}

	t.Run(methodNameJoinToList, func(t *testing.T) {
		l := list.NewList()
		array := NewArray(testCase)
		err := array.JoinToList(l)
		actual := l.String()
		assert.Equal(t, expectedJointToEmptyList, actual)
		assert.NoError(t, err)
	})
}

func TestArray_JoinToEmptyList_Error(t *testing.T) {
	var testCase []int

	t.Run(methodNameJoinToList, func(t *testing.T) {
		l := list.NewList()
		array := NewArray(testCase)
		err := array.JoinToList(l)
		assert.Error(t, err)
	})
}

func TestArray_MoveToArrayError(t *testing.T) {
	testCase := []int{404, 0, 1}

	t.Run(methodNameMoveToArray, func(t *testing.T) {
		s := stack.NewStack()
		array := NewArray(testCase)
		counter, err := array.MoveToArray(s)
		assert.Equal(t, 0, counter)
		assert.Error(t, err)
	})
}

func TestArray_MoveToArray_ArrayOverflow(t *testing.T) {
	testCase := make([]int, 0, 3)

	t.Run(methodNameMoveToArray, func(t *testing.T) {
		s := stack.NewStack()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		s.Push(4)
		s.Push(5)
		array := NewArray(testCase)
		counter, err := array.MoveToArray(s)
		assert.Equal(t, 3, counter)
		assert.NoError(t, err)
	})
}

func TestArray_MoveToArray_StackIsOver(t *testing.T) {
	testCase := make([]int, 0, 10)

	t.Run(methodNameMoveToArray, func(t *testing.T) {
		s := stack.NewStack()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		s.Push(4)
		s.Push(5)
		array := NewArray(testCase)
		counter, err := array.MoveToArray(s)
		assert.Equal(t, 5, counter)
		assert.NoError(t, err)
	})
}

func TestArray_InQueue_Error(t *testing.T) {
	var testCase []int

	t.Run(methodNameInQueue, func(t *testing.T) {
		q := queue.NewQueue()
		q.Push(1)
		array := NewArray(testCase)
		ok, err := array.InQueue(q)
		assert.Equal(t, false, ok)
		assert.Error(t, err)
	})
}

func TestArray_InQueue_True(t *testing.T) {
	testCase := []int{1, 3, 5}

	t.Run(methodNameInQueue, func(t *testing.T) {
		q := queue.NewQueue()
		q.Push(1)
		q.Push(2)
		q.Push(3)
		q.Push(4)
		q.Push(5)
		array := NewArray(testCase)
		ok, err := array.InQueue(q)
		assert.Equal(t, true, ok)
		assert.NoError(t, err)
	})
}

func TestArray_InQueue_False(t *testing.T) {
	testCase := []int{1, 3, 7}

	t.Run(methodNameInQueue, func(t *testing.T) {
		q := queue.NewQueue()
		q.Push(1)
		q.Push(2)
		q.Push(3)
		q.Push(4)
		q.Push(5)
		array := NewArray(testCase)
		ok, err := array.InQueue(q)
		assert.Equal(t, false, ok)
		assert.NoError(t, err)
	})
}
