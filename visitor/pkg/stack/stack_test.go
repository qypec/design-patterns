package stack

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	methodNamePush         = "Push"
	methodNamePop          = "Pop"
	methodNameLen          = "Len"
	notEqualStacksErrorMsg = "Not equal stacks"
)

func TestStack_Push(t *testing.T) {
	expected := &stack{
		head: &elem{
			value: 2,
			prev: &elem{
				value: 1,
				prev:  nil,
			}},
		len: 2,
	}

	t.Run(methodNamePush, func(t *testing.T) {
		s := NewStack()
		s.Push(1)
		s.Push(2)
		if !reflect.DeepEqual(s, expected) {
			t.Error(notEqualStacksErrorMsg, s, expected)
		}
	})
}

func TestStack_Pop(t *testing.T) {
	var value int
	t.Run(methodNamePop, func(t *testing.T) {
		s := NewStack()
		s.Push(1)
		value, _ = s.Pop()
		assert.Equal(t, 1, value)
		s.Push(2)
		s.Push(3)
		value, _ = s.Pop()
		assert.Equal(t, 3, value)
		value, _ = s.Pop()
		assert.Equal(t, 2, value)
	})
}

func TestStack_Len(t *testing.T) {
	t.Run(methodNameLen, func(t *testing.T) {
		s := NewStack()
		s.Push(1)
		assert.Equal(t, 1, s.Len())
		s.Push(2)
		s.Push(3)
		assert.Equal(t, 3, s.Len())
		s.Pop()
		assert.Equal(t, 2, s.Len())
	})
}
