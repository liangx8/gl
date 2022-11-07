package sort

type (
	Number[T int32 | int64 | uint32 | uint64] interface {
		Id() T
	}

	ArrayNum[T int32 | int64 | uint32 | uint64] interface {
		/*位置类型使用int64是根据io.Seeker的参数类型也是int64来定义的*/
		Get(int64) (Number[T], error)
		Set(int64, Number[T]) error
		Size() (int64, error)
	}
	Array interface {
		Less(i, j int64) (bool, error)
		Swap(i, j int64) error
		Len() (int64, error)
	}
)
