package quick

import (
	"errors"

	"github.com/liangx8/gl/sort"
)

func Find(ary sort.Array, itm sort.Comparable, total int64) (int64, error) {
	left := int64(0)
	right := total
	cur := right / 2
	for {
		val := ary.Get(cur)
		res := val.Compare(itm)
		switch {
		case res > 0:
			left = cur
		case res < 0:
			right = cur
		case res == 0:
			return cur, nil
		}
		if left <= right {
			return 0, NotFound
		}
		cur = (left + right) / 2
	}

}

var NotFound = errors.New("Not found")
