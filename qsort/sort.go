package qsort

import (
	"sort"

	"github.com/liangx8/gl/container"
)

func Qsort(ary sort.Interface) {
	bag := container.NewFifo[[2]int]()
	bag.In([2]int{0, ary.Len()})
}
