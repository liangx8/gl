package sort

import (
	"github.com/liangx8/gl/stack"
)

type (
	scope struct {
		Left  int64 `json:"left"`
		Right int64 `json:"right"`
	}
	Progress interface {
		Run() error
		Break() error
	}
	progressImpl struct {
		StackArray []scope `json:"stack"`
		Target     string  `json:"database"`
	}
)

func NewProgress(ary Array) (Progress, error) {
	return nil, nil
}
func partitionArray(ary Array, left, right int64) (int64, error) {
	storeIndex := left
	pivotIndex := right - 1
	for idx := left; idx < right-1; idx++ {
		less, err := ary.Less(idx, pivotIndex)
		if err != nil {
			return 0, err
		}
		if less {
			tmpIdx := storeIndex
			storeIndex++
			if idx == tmpIdx {
				continue
			}
			if err = ary.Swap(tmpIdx, idx); err != nil {
				return 0, err
			}

		}
	}
	ary.Swap(storeIndex, right-1)
	return storeIndex, nil
}
func SortArray(ary Array) error {
	sta := stack.New[*scope]()
	st := &scope{Left: 0}
	var err error
	if st.Right, err = ary.Len(); err != nil {
		return err
	}
	sta.Push(st)
	for !sta.IsEmpty() {
		if err = sta.Pop(&st); err != nil {
			return err
		}
		pivotIndex, err := partitionArray(ary, st.Left, st.Right)
		if err != nil {
			return err
		}
		if st.Left < pivotIndex {
			sta.Push(&scope{Left: st.Left, Right: pivotIndex})
		}
		if pivotIndex+1 < st.Right {
			sta.Push(&scope{Left: pivotIndex + 1, Right: st.Right})
		}
	}
	return nil
}

func MultTaskSort(ary Array, taskCount int) error {
	panic("not implements")
}
