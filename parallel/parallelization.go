// The parallel package provides utilities for running tasks
// concurrently.
package parallel

// Run represents a number of functions running concurrently.
type (
	Run struct {
		work   chan int
		jobcnt int32
	}
)

// NewRun returns a new parallel instance. It provides a way of running
// functions concurrently while limiting the maximum number running at
// once to max. The creater is responsble to monitor all progress done
func NewRun(max int) *Run {
	if max < 1 {
		panic("parameter max must be >= 1")
	}
	return &Run{
		work: make(chan int, max),
	}
}

// Do requests that r run f concurrently.  If there are already the maximum
// number of functions running concurrently, it will block until one of them
// has completed.
func (r *Run) Do(f func()) {
	r.work <- 1
	go func(fn func()) {
		fn()
		<-r.work
	}(f)
}
