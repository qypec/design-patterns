package queue

type elem struct {
	value int
	next  *elem
}

func newElem(value int) *elem {
	return &elem{
		value: value,
		next:  nil,
	}
}
