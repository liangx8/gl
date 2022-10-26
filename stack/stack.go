package stack

import "errors"

type (
	node[T any] struct {
		data T
		next *node[T]
	}
	Stack[T any] struct {
		top *node[T]
	}
)

func New[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (堆栈 *Stack[T]) Push(t T) {
	n := &node[T]{data: t, next: 堆栈.top}
	堆栈.top = n
}
func (ss *Stack[T]) Pop(对象 *T) error {
	if ss.top == nil {
		return errEmpty
	}
	ret := ss.top
	ss.top = ss.top.next
	*对象 = ret.data
	return nil
}
func (ss *Stack[T]) IsEmpty() bool {
	return ss.top == nil
}

var errEmpty = errors.New("Stack is empty")
