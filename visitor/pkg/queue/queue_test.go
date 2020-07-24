package queue

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	methodNamePush   = "Push"
	methodNamePop    = "Pop"
	methodNameLen    = "Len"
	methodNameSearch = "Search"
	notEqualQueues   = "Not equal queues"
)

func TestQueue_Push(t *testing.T) {
	expected := &queue{
		head: &elem{
			value: 1,
			next: &elem{
				value: 2,
				next:  nil,
			}},
		len: 2,
	}

	t.Run(methodNamePush, func(t *testing.T) {
		q := NewQueue()
		q.Push(1)
		q.Push(2)
		if !reflect.DeepEqual(q, expected) {
			t.Error(notEqualQueues, q, expected)
		}
	})
}

func TestQueue_Pop(t *testing.T) {
	t.Run(methodNamePop, func(t *testing.T) {
		q := NewQueue()
		q.Push(1)
		assert.Equal(t, 1, q.Pop())
		q.Push(2)
		q.Push(3)
		assert.Equal(t, 2, q.Pop())
		assert.Equal(t, 3, q.Pop())
	})
}

func TestQueue_Len(t *testing.T) {
	t.Run(methodNameLen, func(t *testing.T) {
		q := NewQueue()
		q.Push(1)
		assert.Equal(t, 1, q.Len())
		q.Push(2)
		q.Push(3)
		assert.Equal(t, 3, q.Len())
		q.Pop()
		assert.Equal(t, 2, q.Len())
	})
}

func TestQueue_Search(t *testing.T) {
	t.Run(methodNameSearch, func(t *testing.T) {
		q := NewQueue()
		assert.Equal(t, false, q.Search(1))
		q.Push(1)
		assert.Equal(t, true, q.Search(1))
		assert.Equal(t, false, q.Search(100))
	})
}
