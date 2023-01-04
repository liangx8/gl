package bubble

import "github.com/liangx8/gl/sort"

func Sort(ary sort.Array) {
	swapped := true
	top := ary.Len() - 1
	for swapped {
		swapped = false
		for ii := int64(0); ii < top; ii++ {
			itm1 := ary.Get(ii)
			itm2 := ary.Get(ii + 1)
			if itm1.Compare(itm2) >= 0 {
				ary.Set(ii, itm2)
				ary.Set(ii+1, itm1)
				swapped = true
			}
		}
		top = top - 1
	}
}
