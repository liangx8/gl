package quick_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/liangx8/gl/sort/quick"
)

type (
	aryImpl struct {
		data []int
	}
	daoImpl struct {
		ary          *aryImpl
		store, pivot int64
	}
)

func (ary *aryImpl) Part(pivotIndex int64) (quick.ArrayIO, error) {
	dao := &daoImpl{ary: ary, pivot: pivotIndex}
	return dao, nil
}
func (ary *aryImpl) Len() (int64, error) {
	return int64(len(ary.data)), nil
}
func (ary *aryImpl) String() string {
	return fmt.Sprintf("%v", ary.data)
}

func (dao *daoImpl) LessPivot(idx int64) (bool, error) {

	return dao.ary.data[idx] < dao.ary.data[dao.pivot], nil
}

func (dao *daoImpl) Store(idx int64) error {
	dao.store = idx
	return nil
}
func (dao *daoImpl) SwapStore(idx int64) error {
	dao.ary.data[idx], dao.ary.data[dao.store] = dao.ary.data[dao.store], dao.ary.data[idx]
	return nil
}
func (dao *daoImpl) SwapStorePivot() error {
	dao.ary.data[dao.pivot], dao.ary.data[dao.store] = dao.ary.data[dao.store], dao.ary.data[dao.pivot]
	return nil
}

func createTestAry(cnt int) *aryImpl {
	ar := &aryImpl{data: make([]int, cnt)}
	for ix := 0; ix < cnt; ix++ {
		ar.data[ix] = rand.Int()
	}
	return ar
}

const num = 10000

func TestSortArray(t *testing.T) {
	ary := createTestAry(num)
	quick.Sort(ary)
	for ix := int64(0); ix < num-1; ix++ {
		if ary.data[ix] > ary.data[ix+1] {
			t.Fail()
		}
	}
}
