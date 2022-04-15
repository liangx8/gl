package stack_test

import (
	"testing"

	"github.com/liangx8/gl/stack"
)

func TestStack(t *testing.T) {
	ss := stack.New[int]()
	var v int
	ss.Push(1)
	ss.Push(2)
	if ss.Pop(&v); v != 2 {
		t.Fatal("Stack test fail")
	}
	if ss.Pop(&v); v != 1 {
		t.Fatal("Stack test fail")
	}
	if err := ss.Pop(&v); err == nil {
		t.Fatal("Stack test fail")
	}
}
