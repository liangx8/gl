package sort_test

import (
	"math/rand"
	"testing"

	"github.com/liangx8/gl/sort"
	"github.com/liangx8/gl/sort/bubble"
	"github.com/liangx8/gl/sort/quick"
)

type (
	ele    int
	eleAry []int
)

func (e ele) Compare(r sort.Comparable) int {
	return int(e - r.(ele))
}
func (ea eleAry) Get(x int64) sort.Comparable {
	return ele(ea[x])
}
func (ea eleAry) Set(x int64, c sort.Comparable) {
	ea[x] = int(c.(ele))
}
func (ea eleAry) Len() int64 {
	return int64(len(ea))
}
func createArray(size int64) sort.Array {
	ia := make([]int, size)
	for idx := range ia {
		ia[idx] = rand.Int() % 100
	}
	return eleAry(ia)
}

func TestBubbleSort(t *testing.T) {
	const max int64 = 20000
	ary := createArray(max)
	bubble.Sort(ary)
	if !sort.Order(ary) {
		t.Fail()
	}
}
func TestQuickSort(t *testing.T) {
	const max int64 = 90000
	ary := createArray(max)
	quick.Sort(ary)
	if !sort.Order(ary) {
		t.Fail()
	}

}

func BenchmarkBubbleSort(b *testing.B) {
	const max int64 = 1000
	ary := createArray(max)
	b.ResetTimer()
	for t := 0; t < b.N; t++ {
		bubble.Sort(ary)
	}

}
func BenchmarkQuickSort(b *testing.B) {
	// BenchmarkQuickSort-4   	       1	1504717145 ns/op	 5689328 B/op	    1878 allocs/op
	const max int64 = 180000
	ary := createArray(max)
	b.ResetTimer()
	for t := 0; t < b.N; t++ {
		quick.Sort(ary)
	}

}
