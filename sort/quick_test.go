package sort_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/liangx8/gl/sort"
)

type (
	entityImpl struct {
		num uint32
	}
	arrayImpl struct {
		data []sort.Number[uint32]
	}
	array1Impl []int
)

func (impl *entityImpl) Id() uint32 {
	return impl.num
}
func (impl *entityImpl) String() string {
	return fmt.Sprint(impl.num)
}
func (impl *arrayImpl) Get(pos int64) (sort.Number[uint32], error) {
	return impl.data[pos], nil
}
func (impl *arrayImpl) Set(pos int64, ent sort.Number[uint32]) error {
	impl.data[pos] = ent
	return nil
}
func (impl *arrayImpl) Size() (int64, error) {
	return int64(len(impl.data)), nil
}
func (impl *arrayImpl) String() string {
	return fmt.Sprint(impl.data)
}
func createArray(num int64) *arrayImpl {
	ary := &arrayImpl{data: make([]sort.Number[uint32], num)}
	for ix := int64(0); ix < num; ix++ {
		ary.Set(ix, &entityImpl{num: rand.Uint32() % 100})
	}
	return ary
}

func TestQuickSort(t *testing.T) {
	// prepare data
	const size = int64(10)
	ary := createArray(size)
	t.Log(ary)
	if err := sort.Sort[uint32](ary); err != nil {
		t.Fatal(err)
	}
	t.Log(ary)
	for ix := int64(0); ix < size-1; ix++ {
		v1, _ := ary.Get(ix)
		v2, _ := ary.Get(ix + 1)
		if v1.Id() > v2.Id() {
			t.Log(ix, v1.Id())
			t.FailNow()
		}
	}
}
func TestQuickNotRecur(t *testing.T) {
	const size = int64(10)
	ary := createArray(size)
	t.Log(ary)
	if err := sort.QuickNotRecur[uint32](ary); err != nil {
		t.Fatal(err)
	}
	for ix := int64(0); ix < size-1; ix++ {
		v1, _ := ary.Get(ix)
		v2, _ := ary.Get(ix + 1)
		if v1.Id() > v2.Id() {
			t.Log(ix, v1.Id())
			t.FailNow()
		}
	}
}

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
