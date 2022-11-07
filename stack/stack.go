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
func CreateFromArray[T any](ary []T) *Stack[T] {
	st := &Stack[T]{nil}
	for _, da := range ary {
		n := &node[T]{data: da, next: st.top}
		st.top = n
	}
	return st
}

func (堆栈 *Stack[T]) Push(t T) {
	n := &node[T]{data: t, next: 堆栈.top}
	堆栈.top = n
}
func (ss *Stack[T]) Pop(对象 *T) error {
	if ss.top == nil {
		return errors.New("Stack is empty")
	}
	ret := ss.top
	ss.top = ss.top.next
	*对象 = ret.data
	return nil
}
func (ss *Stack[T]) IsEmpty() bool {
	return ss.top == nil
}
func (ss *Stack[T]) ToArray() []T {
	ary := make([]T, 0, 40)
	top := ss.top
	for top != nil {
		ary = append(ary, top.data)
		top = top.next
	}
	return ary
}
