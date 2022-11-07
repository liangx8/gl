package stack_test

import (
	"testing"

	"github.com/liangx8/gl/stack"
)

type (
	Int struct {
		Val int
	}
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

func TestObject(t *testing.T) {
	v1 := Int{1}
	ss := stack.New[Int]()
	if !ss.IsEmpty() {
		t.Fail()
	}
	v2 := Int{2}
	ss.Push(v1)
	ss.Push(v2)
	if ss.IsEmpty() {
		t.Fail()
	}
	var v3 Int
	if ss.Pop(&v3); v3 != v2 {
		t.Fail()
	}
	if ss.Pop(&v3); v3 != v1 {
		t.Fail()
	}
	if !ss.IsEmpty() {
		t.Fail()
	}
	if err := ss.Pop(&v3); err == nil {
		t.Fatal("Stack test fail")
	}
}
func TestStackRestore(t *testing.T) {
	s1 := stack.New[int]()
	s1.Push(1)
	s1.Push(2)
	s1ary := s1.ToArray()
	s2 := stack.CreateFromArray(s1ary)
	var val int
	s2.Pop(&val)
	if val != 2 {
		t.Fatalf("expect 2 but %d", val)
	}
	s2.Pop(&val)
	if val != 1 {
		t.Fatalf("expect 1 but %d", val)
	}
}
