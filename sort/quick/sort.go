// Disk IO is one of critical issues for sorting a big database
// algorithm reference: https://en.wikipedia.org/wiki/Quicksort

package quick

import (
	"github.com/liangx8/gl/stack"
)

type (
	scope struct {
		Left  int64 `json:"left"`
		Right int64 `json:"right"`
	}
	Comparable interface {
		Compare(Comparable) int
	}
	//
	Array interface {
		Len() (int64, error)
		Load(idx int64) (Comparable, error)
		Set(idx int64, com Comparable) error
	}
)

func partition(ary Array, left, right int64) (int64, error) {
	pivotIdx := right - 1
	pivot, err := ary.Load(pivotIdx)
	if err != nil {
		return 0, err
	}
	storeIdx := left
	store, err := ary.Load(storeIdx)
	if err != nil {
		return 0, err
	}
	for idx := left; idx < right-1; idx++ {

		cur, err := ary.Load(idx)
		if err != nil {
			return 0, err
		}
		if cur.Compare(pivot) < 0 {
			if idx != storeIdx {
				if err = ary.Set(idx, store); err != nil {
					return 0, err
				}
				if err = ary.Set(storeIdx, cur); err != nil {
					return 0, err
				}

			}
			storeIdx++
			if store, err = ary.Load(storeIdx); err != nil {
				return 0, err
			}
		}

	}
	if err = ary.Set(storeIdx, pivot); err != nil {
		return 0, err
	}
	if err = ary.Set(pivotIdx, store); err != nil {
		return 0, err
	}
	return storeIdx, nil

}
func Sort(ary Array) error {
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
		pivotIndex, err := partition(ary, st.Left, st.Right)
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
