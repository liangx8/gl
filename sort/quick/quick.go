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
func sort3(ary sort.Array, start int64) {
	v1 := ary.Get(start)
	v2 := ary.Get(start + 1)
	v3 := ary.Get(start + 2)
	if v1.Compare(v2) > 0 {
		if v1.Compare(v3) > 0 {
			ary.Set(start+2, v1)
			if v2.Compare(v3) > 0 {
				ary.Set(start, v3)
				ary.Set(start+1, v2)
			} else {
				ary.Set(start, v2)
				ary.Set(start+1, v3)

			}
		} else {
			ary.Set(start, v2)
			ary.Set(start+1, v1)
			ary.Set(start+2, v3)
		}
		return
	}
	// TODO: else
}
func Sort(ary sort.Array) {
	qu := container.NewFifo[[2]int64]()
	qu.In([2]int64{0, ary.Len()})
	for qu.Len() > 0 {
		scope := qu.Out()
		idx := partition(ary, scope)
		switch idx - scope[0] {
		case 0, 1:
		case 2:
			v1 := ary.Get(scope[0])
			v2 := ary.Get(idx - 1)
			if v1.Compare(v2) > 0 {
				ary.Set(scope[0], v2)
				ary.Set(idx-1, v1)
			}
		case 3:
			sort3(ary, scope[0])
		default:
			qu.In([2]int64{scope[0], idx})
		}
		switch scope[1] - idx - 1 {
		case 0, 1:
		case 2:
			v1 := ary.Get(idx + 1)
			v2 := ary.Get(scope[1] - 1)
			if v1.Compare(v2) > 0 {
				ary.Set(idx+1, v2)
				ary.Set(scope[1]-1, v1)
			}
		case 3:
			sort3(ary, idx+1)
		default:
			qu.In([2]int64{idx + 1, scope[1]})
		}
	}

}
