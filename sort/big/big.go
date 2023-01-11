/*
对大的数据进行排序，主要场景是大文件。因此为对象设计了closer接口，以便于数据处理结束时的持久化
*/
package big

import (
	"errors"
	"io"

	"github.com/liangx8/gl/container"
	"github.com/liangx8/gl/sort"
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
	SORT_STATUS int = iota
	SORT_BREAK
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

func MultQuickSort(ab ArrayBuilder, bi *BreakInfo, ask <-chan int, ans chan<- any) {
	fifo := container.NewFifoWith(bi.Parts)
	jobCnt := fifo.Len()
	job := make(chan [2]int64)
	results := make(chan [3]int64) // 顺序 [left,right,pivot]

	occupy := true
	current := fifo.Out()
	for jobCnt > 0 {
		if occupy {
			select {
			case job <- current:
				occupy = false
			}
		} else {
			select {
			case a := <-ask:
				switch a {
				case SORT_STATUS:
					ans <- "return status"
				case SORT_BREAK:
					ans <- SORT_DONE
					return
				}
			case r := <-results:
				if r[0]+1 < r[2] {

				}
			}
		}
	}
}
