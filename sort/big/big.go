/*
对大的数据进行排序，主要场景是大文件。因此为对象设计了closer接口，以便于数据处理结束时的持久化
*/
package big

import (
	"errors"
	"io"
	"runtime"

	"github.com/liangx8/gl/container"
	"github.com/liangx8/gl/parallel"
	"github.com/liangx8/gl/sort"
	"github.com/liangx8/gl/sort/quick"
)

type (
	ArrayCloser interface {
		sort.Array
		io.Closer
	}
	ArrayBuilder interface {
		// 把磁盘中的内容整块读入内存进行处理
		Buff(offset int64, limit int) (ArrayCloser, error)
		Unbuff() (ArrayCloser, error)
	}
	BreakInfo struct {
		Parts [][2]int64
	}
)

const (
	REQUEST_STATUS int = iota
	REQUEST_BREAK
)

var SORT_DONE = errors.New("Sort Done")

// 对于大数据的排序，需要很长时间，因此需要查询处理的进度。
func partition(ac ArrayCloser, scope [2]int64, ans chan<- any) int64 {
	/*
	   left := scope[0]
	   right := scope[1]
	   pivot := ac.Get(right - 1)
	*/
	return 0
}
func service(ab ArrayBuilder, ans chan<- any, chIn <-chan [2]int64, chOut chan<- any) {
	run := parallel.NewRun(runtime.NumCPU())

	for jo := range chIn {
		jo := jo
		run.Do(func() {
			ary, err := ab.Unbuff()
			if err != nil {
				ans <- err
				return
			}
			x := partition(ary, jo, ans)
			err = ary.Close()
			if err != nil {
				ans <- err
				return
			}
			// 检查结果x 使用不同的排序策略
			size := x - jo[0]
			if size > 3000 {
				chOut <- [2]int64{jo[0], x}
			} else {
				ary, err = ab.Buff(jo[0], int(size))
				if err != nil {
					ans <- err
					return
				}
				quick.Sort(ary)
				err = ary.Close()
				if err != nil {
					ans <- err
					return
				}
			}
			size = jo[1] - x - 1
			if size > 3000 {
				chOut <- [2]int64{x + 1, jo[1]}
			} else {
				ary, err = ab.Buff(x+1, int(size))
				if err != nil {
					ans <- err
					return
				}
				quick.Sort(ary)
				err = ary.Close()
				if err != nil {
					ans <- err
					return
				}
			}
			chOut <- 1 // current job is done
		})
	}
	close(chOut)
}
func MultQuickSort(ab ArrayBuilder, bi *BreakInfo, ask <-chan int, ans chan<- any) {
	fifo := container.NewFifoWith(bi.Parts)
	jobCnt := fifo.Len()
	job := make(chan [2]int64)
	results := make(chan any)
	go service(ab, ans, job, results)

	occupy := true
	current := fifo.Out()

	for jobCnt > 0 {
		if occupy {
			select {
			case job <- current:
				if fifo.Len() > 0 {
					current = fifo.Out()
				} else {
					occupy = false
				}
			}
		} else {
			select {
			case a := <-ask:
				switch a {
				case REQUEST_STATUS:
					ans <- "return status"
				case REQUEST_BREAK:
					ans <- SORT_DONE
					return
				}
			case r := <-results:
				switch val := r.(type) {
				case int:
					jobCnt--
				case [2]int64:
					fifo.In(val)
					jobCnt++
				default:
					panic("impossible")
				}
			}
		}
	}
	close(job)
	<-results
}
