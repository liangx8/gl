package sort

import (
	"github.com/liangx8/gl/stack"
)

type (
	Status struct {
		Left, Right int64
	}
)

func partitionNum[T int32 | int64 | uint32 | uint64](ary ArrayNum[T], left, right int64) (int64, error) {
	pivotValue, err := ary.Get(right - 1)
	if err != nil {
		return 0, err
	}
	storeIndex := left
	for i := left; i < right-1; i++ {
		idata, err := ary.Get(i)
		if err != nil {
			return 0, err
		}
		if idata.Id() < pivotValue.Id() {
			// swap
			tmpIdx := storeIndex
			storeIndex++
			if i == tmpIdx {
				continue
			}
			storeValue, err := ary.Get(tmpIdx)
			if err != nil {
				return 0, err
			}
			if err := ary.Set(tmpIdx, idata); err != nil {
				return 0, err
			}
			if err := ary.Set(i, storeValue); err != nil {
				return 0, err
			}
		}
	}
	storeValue, err := ary.Get(storeIndex)
	if err != nil {
		return 0, err
	}
	if err := ary.Set(right-1, storeValue); err != nil {
		return 0, err
	}
	if err := ary.Set(storeIndex, pivotValue); err != nil {
		return 0, err
	}
	return storeIndex, nil

}

func quickSort[T int32 | int64 | uint32 | uint64](ary ArrayNum[T], left, right int64) error {
	if left < right {

		pivotIndex, err := partitionNum(ary, left, right)
		if err != nil {
			return err
		}
		if err = quickSort(ary, left, pivotIndex); err != nil {
			return err
		}
		if err = quickSort(ary, pivotIndex+1, right); err != nil {
			return err
		}
	}
	return nil
}
func QuickNotRecur[T int32 | int64 | uint32 | uint64](ary ArrayNum[T]) error {
	sta := stack.New[*Status]()
	st := &Status{Left: 0}
	var err error
	if st.Right, err = ary.Size(); err != nil {
		return err
	}
	sta.Push(st)
	for !sta.IsEmpty() {
		if err = sta.Pop(&st); err != nil {
			return err
		}
		pivotIndex, err := partitionNum(ary, st.Left, st.Right)
		if err != nil {
			return err
		}
		if st.Left < pivotIndex {
			sta.Push(&Status{Left: st.Left, Right: pivotIndex})
		}
		if pivotIndex+1 < st.Right {
			sta.Push(&Status{Left: pivotIndex + 1, Right: st.Right})
		}
	}
	return nil
}
func Sort[T int32 | int64 | uint32 | uint64](ary ArrayNum[T]) error {
	end, err := ary.Size()
	if err != nil {
		return err
	}
	return quickSort(ary, 0, end)
}
