package rutils

import "math/rand"

var extchar = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '.', ',', '-', '+', '/'}

func RandomString(size int) string {
	ct := make([]rune, size)
	src := int64(0)
	cnt := int64(len(extchar))
	for ix := range ct {
		if src == 0 {
			src = rand.Int63()
		}
		mod := src % cnt
		src = src / cnt
		ct[ix] = extchar[mod]
	}
	return string(ct)
}
