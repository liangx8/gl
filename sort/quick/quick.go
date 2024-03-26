package quick

import (
	"github.com/liangx8/gl/container"
	"github.com/liangx8/gl/sort"
)

func partition(ary sort.Array, scope [2]int64) int64 {
	left := scope[0]
	right := scope[1] - 1
	pivot := ary.Get(right)
	storeIndex := left
	store := ary.Get(storeIndex)
	for store.Compare(pivot) < 0 {
		storeIndex++
		store = ary.Get(storeIndex)
		if storeIndex == right {
			return storeIndex
		}
	}
	for idx := storeIndex + 1; idx < right; idx++ {
		cur := ary.Get(idx)
		if cur.Compare(pivot) < 0 {
			ary.Set(storeIndex, cur)
			ary.Set(idx, store)
			storeIndex++
			store = ary.Get(storeIndex)
		}
	}
	ary.Set(storeIndex, pivot)
	ary.Set(right, store)
	return storeIndex
}
func sort3(ary sort.Array, start int64, size int64) {
	swapped := true
	top := start + size - 1
	for swapped {
		swapped = false
		for idx := start; idx < top; idx++ {
			v1 := ary.Get(idx)
			v2 := ary.Get(idx + 1)
			if v1.Compare(v2) > 0 {
				ary.Set(idx, v2)
				ary.Set(idx+1, v1)
				swapped = true
			}
		}
		top = top - 1
	}
}
func Sort(ary sort.Array) {
	qu := container.NewFifo[[2]int64]()
	qu.In([2]int64{0, ary.Len()})
	for qu.Len() > 0 {
		scope := qu.Out()
		idx := partition(ary, scope)
		wd := idx - scope[0]
		switch wd {
		case 0, 1:
		case 2:
			v1 := ary.Get(scope[0])
			v2 := ary.Get(idx - 1)
			if v1.Compare(v2) > 0 {
				ary.Set(scope[0], v2)
				ary.Set(idx-1, v1)
			}
		default:
			if wd < 10 {
				sort3(ary, scope[0], wd)
			} else {
				qu.In([2]int64{scope[0], idx})
			}
		}
		wd = scope[1] - idx - 1
		switch wd {
		case 0, 1:
		case 2:
			v1 := ary.Get(idx + 1)
			v2 := ary.Get(scope[1] - 1)
			if v1.Compare(v2) > 0 {
				ary.Set(idx+1, v2)
				ary.Set(scope[1]-1, v1)
			}
		default:
			if wd < 10 {
				sort3(ary, idx+1, wd)
			} else {
				qu.In([2]int64{idx + 1, scope[1]})
			}
		}
	}

}
