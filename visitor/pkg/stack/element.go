package stack

type elem struct {
	value int
	prev  *elem
}

func newElem(value int) elem {
	return elem{
		value: value,
	}
}
