// The parallel package provides utilities for running tasks
// concurrently.
package parallel

// Run represents a number of functions running concurrently.
type (
	Run struct {
		work       chan int
		res        chan<- error
		Errs       <-chan error
		atomicUint uint32
	}
)

// NewRun returns a new parallel instance. It provides a way of running
// functions concurrently while limiting the maximum number running at
// once to max. The creater is responsble to monitor all progress done
func NewRun(max int) *Run {
	if max < 1 {
		panic("parameter max must be >= 1")
	}
	result := make(chan error)
	return &Run{
		work:       make(chan int, max),
		res:        result,
		Errs:       result,
		atomicUint: 0,
	}
}

// Do requests that r run f concurrently.  If there are already the maximum
// number of functions running concurrently, it will block until one of them
// has completed.
func (r *Run) Do(f func() error) {
	r.atomicUint++
	r.work <- 1
	go func(fn func() error, res chan<- error) {
		if err := fn(); err != nil {
			res <- err
		}
		<-r.work
		r.atomicUint--
	}(f, r.res)
}
