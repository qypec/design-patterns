package list

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
	expected := &list{
		head: &elem{
			value: 2,
			next: &elem{
				value: 1,
				next:  nil,
			}},
		len: 2,
	}

	t.Run(methodNamePush, func(t *testing.T) {
		l := NewList()
		l.Push(1)
		l.Push(2)
		if !reflect.DeepEqual(l, expected) {
			t.Error("Not equal lists\n", l, expected)
		}
	})
}

func TestList_Pop(t *testing.T) {
	t.Run(methodNamePop, func(t *testing.T) {
		l := NewList()
		l.Push(1)
		assert.Equal(t, 1, l.Pop())
		l.Push(2)
		l.Push(3)
		assert.Equal(t, 3, l.Pop())
		assert.Equal(t, 2, l.Pop())
	})
}

func TestList_Len(t *testing.T) {
	t.Run(methodNameLen, func(t *testing.T) {
		l := NewList()
		l.Push(1)
		assert.Equal(t, 1, l.Len())
		l.Push(2)
		l.Push(3)
		assert.Equal(t, 3, l.Len())
		l.Pop()
		assert.Equal(t, 2, l.Len())
	})
}
