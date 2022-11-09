package sort

type (
	Array interface {
		Less(i, j int64) (bool, error)
		Swap(i, j int64) error
		Len() (int64, error)
	}
)
