package stack

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	methodNamePush = "Push"
	methodNamePop  = "Pop"
	methodNameLen  = "Len"
)

func TestList_Push(t *testing.T) {
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
			t.Error("Not equal stacks\n", s, expected)
		}
	})
}

func TestList_Pop(t *testing.T) {
	t.Run(methodNamePop, func(t *testing.T) {
		s := NewStack()
		s.Push(1)
		assert.Equal(t, 1, s.Pop())
		s.Push(2)
		s.Push(3)
		assert.Equal(t, 3, s.Pop())
		assert.Equal(t, 2, s.Pop())
	})
}

func TestList_Len(t *testing.T) {
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
