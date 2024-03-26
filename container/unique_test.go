package container_test

import (
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"testing"
)

var iary = []int{0, 1, 2, 3, 5, 5, 6, 7, 8, 9, 10}
var sary = []string{"one", "two", "three", "four", "five"}

func createArray(size int64) []int {
	ia := make([]int, size)
	for idx := range ia {
		ia[idx] = rand.Int()
	}
	return ia
}

func BenchmarkUnique(b *testing.B) {
	ia := createArray(1000000)
	sa := sort.IntSlice(ia)
	sa.Sort()
	b.ResetTimer()
	if !sort.IntsAreSorted(sa) {
		b.Fail()
	}
}
func TestXX(t *testing.T) {
	pos := sort.SearchInts(iary, 4)
	if pos != 4 {
		t.Fail()
	}
	pos = sort.SearchInts(iary, 5)
	if pos != 4 {
		t.Fail()
	}
	pos = sort.SearchInts(iary, 6)
	if pos != 6 {
		t.Fail()
	}
}
func TestSomething(t *testing.T) {
	one := "one"
	one1 := one
	clo := strings.Clone(one)
	t.Logf("%p,%p", &sary[0], &one)
	if !reflect.DeepEqual(one, one1) {
		t.Fail()
	}
	if !reflect.DeepEqual(one, sary[0]) {
		t.Fail()
	}
	if !reflect.DeepEqual(one, clo) {
		t.Fail()
	}
}
