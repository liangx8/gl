package sort_test

import (
	"math/rand"
	"testing"

	"github.com/liangx8/gl/sort"
)

type (
	array1Impl []int
)

func (ary array1Impl) Less(i, j int64) (bool, error) {
	return ary[i] < ary[j], nil
}
func (ary array1Impl) Swap(i, j int64) error {
	ary[i], ary[j] = ary[j], ary[i]
	return nil
}
func (ary array1Impl) Len() (int64, error) {
	return int64(len(ary)), nil
}
func TestQuick1(t *testing.T) {
	ary := array1Impl(make([]int, 10))
	for ix := 0; ix < 10; ix++ {
		ary[ix] = rand.Int()
	}
	sort.SortArray(ary)
	size, _ := ary.Len()
	for ix := int64(0); ix < size-1; ix++ {
		if ary[ix] > ary[ix+1] {
			t.Fail()
		}
	}
}
