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
func Sort(ary sort.Array) {
	qu := container.NewFifo[[2]int64]()
	qu.In([2]int64{0, ary.Len()})
	for qu.Len() > 0 {
		scope := qu.Out()
		idx := partition(ary, scope)
		if idx-1 > scope[0] {
			qu.In([2]int64{scope[0], idx})
		}
		if scope[1]-2 > idx {
			qu.In([2]int64{idx + 1, scope[1]})
		}
	}

}
