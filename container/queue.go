package container

type (
	Fifo[T any] struct {
		data []T
	}
)

func (q *Fifo[T]) Len() int {
	return len(q.data)
}
func (q *Fifo[T]) In(e T) {
	q.data = append(q.data, e)
}
func (q *Fifo[T]) Out() T {
	e := q.data[0]
	q.data = q.data[1:]
	return e
}
func (q *Fifo[T]) AsArray() []T {
	return q.data
}
func NewFifo[T any]() *Fifo[T] {
	return &Fifo[T]{data: make([]T, 0)}
}
func NewFifoWith[T any](data []T) *Fifo[T] {
	return &Fifo[T]{data: data}
}
