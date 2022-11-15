package quick_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/liangx8/gl/sort/quick"
)

type (
	cmpImpl int
	aryImpl []quick.Comparable
)

func (cmp cmpImpl) Compare(c quick.Comparable) int {
	return int(cmp - c.(cmpImpl))
}
func (cmp cmpImpl) String() string {
	return fmt.Sprint(int(cmp))
}

func (ary aryImpl) Set(idx int64, com quick.Comparable) error {
	ary[idx] = com
	return nil
}

func (ary aryImpl) Load(idx int64) (quick.Comparable, error) {
	return ary[idx], nil
}
func (ary aryImpl) Len() (int64, error) {
	return int64(len(ary)), nil
}

func createTestAry(cnt int) aryImpl {
	ar := aryImpl(make([]quick.Comparable, cnt))
	for ix := 0; ix < cnt; ix++ {
		ar[ix] = cmpImpl(rand.Int() % 100)
	}
	return ar
}

const num = 10

func TestSortArray(t *testing.T) {
	ary := createTestAry(num)
	t.Log(ary)
	quick.Sort(ary)
	t.Log(ary)
	for ix := int64(0); ix < num-1; ix++ {
		if ary[ix].Compare(ary[ix+1]) > 0 {
			t.Fail()
		}
	}
}
