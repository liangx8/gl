package leakybuffer

type (
	LeakyBuffer[T any] struct {
		freeList chan T
		create   func() T
	}
)

func New[T any](capacity int, cr func() T) *LeakyBuffer[T] {
	var lb LeakyBuffer[T]
	lb.freeList = make(chan T, capacity)
	lb.create = cr
	return &lb
}

func (lb *LeakyBuffer[T]) New() T {
	select {
	case b := <-lb.freeList:
		return b
	default:
		// None free, so allocate a new one.
		return lb.create()
	}
}
func (lb *LeakyBuffer[T]) Free(t T) {
	select {
	case lb.freeList <- t:
		// Buffer on free list; nothing more to do.
	default:
		// Free list full, just carry on.
	}
}
