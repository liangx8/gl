package leakybuffer_test

import (
	"testing"

	"github.com/liangx8/gl/leakybuffer"
)

func cr() []byte {

	return make([]byte, 12)
}
func TestLeakyBuffer(t *testing.T) {
	lb := leakybuffer.New(10, cr)
	obj := lb.New()
	t.Log(obj)
	lb.Free(obj)
	if len(obj) != 12 {
		t.Fail()
	}
}
