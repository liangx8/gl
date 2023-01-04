package sort

type (
	Comparable interface {
		Compare(Comparable) int
	}
	Array interface {
		Get(int64) Comparable
		Set(int64, Comparable)
		Len() int64
	}
)

func Order(ary Array) bool {
	for x := int64(0); x < ary.Len()-1; x++ {
		val1 := ary.Get(x)
		val2 := ary.Get(x + 1)
		if val1.Compare(val2) > 0 {
			return false
		}

	}
	return true
}
