package quick

import "github.com/liangx8/gl/stack"

type (
	scope struct {
		Left  int64 `json:"left"`
		Right int64 `json:"right"`
	}
	ArrayDao interface {
		LessPivot(idx int64) (bool, error)
		Store(idx int64) error
		SwapStore(idx int64) error
		SwapStorePivot() error
	}
	Array interface {
		Part(pivotIndex int64) (ArrayDao, error)
		Len() (int64, error)
	}
)

func partition(ary Array, left, right int64) (int64, error) {
	part, err := ary.Part(right - 1)
	if err != nil {
		return 0, err
	}
	storeIdx := left
	if err = part.Store(left); err != nil {
		return 0, err
	}
	for idx := left; idx < right-1; idx++ {
		less, err := part.LessPivot(idx)
		if err != nil {
			return 0, err
		}
		if less {
			if err = part.SwapStore(idx); err != nil {
				return 0, err
			}
			storeIdx++
			if err = part.Store(storeIdx); err != nil {
				return 0, err
			}
		}

	}
	if err = part.SwapStorePivot(); err != nil {
		return 0, err
	}
	return storeIdx, nil

}
func Sort(ary Array) error {
	sta := stack.New[*scope]()
	st := &scope{Left: 0}
	var err error
	if st.Right, err = ary.Len(); err != nil {
		return err
	}
	sta.Push(st)
	for !sta.IsEmpty() {
		if err = sta.Pop(&st); err != nil {
			return err
		}
		pivotIndex, err := partition(ary, st.Left, st.Right)
		if err != nil {
			return err
		}
		if st.Left < pivotIndex {
			sta.Push(&scope{Left: st.Left, Right: pivotIndex})
		}
		if pivotIndex+1 < st.Right {
			sta.Push(&scope{Left: pivotIndex + 1, Right: st.Right})
		}
	}
	return nil
}
