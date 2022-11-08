package sort

import "github.com/liangx8/gl/stack"

type (
	dataWrap struct {
		pivot, left, right int64
	}
	param struct {
		left, right int64
	}
	result struct {
		left, pivot, right int64
	}
)

func partitionArray(ary Array, left, right int64) (int64, error) {
	storeIndex := left
	pivotIndex := right - 1
	for idx := left; idx < right-1; idx++ {
		less, err := ary.Less(idx, pivotIndex)
		if err != nil {
			return 0, err
		}
		if less {
			tmpIdx := storeIndex
			storeIndex++
			if idx == tmpIdx {
				continue
			}
			if err = ary.Swap(tmpIdx, idx); err != nil {
				return 0, err
			}

		}
	}
	ary.Swap(storeIndex, right-1)
	return storeIndex, nil
}
func SortArray(ary Array) error {
	sta := stack.New[*Status]()
	st := &Status{Left: 0}
	var err error
	if st.Right, err = ary.Len(); err != nil {
		return err
	}
	sta.Push(st)
	for !sta.IsEmpty() {
		if err = sta.Pop(&st); err != nil {
			return err
		}
		pivotIndex, err := partitionArray(ary, st.Left, st.Right)
		if err != nil {
			return err
		}
		if st.Left < pivotIndex {
			sta.Push(&Status{Left: st.Left, Right: pivotIndex})
		}
		if pivotIndex+1 < st.Right {
			sta.Push(&Status{Left: pivotIndex + 1, Right: st.Right})
		}
	}
	return nil
}
func run(ary Array, taskCount int, result chan *dataWrap) error {
	pach := make(chan *param, 0)
	resultch := make(chan *dataWrap, 0)

	errch := make(chan error, 0)
	go func() {
		thr := make(chan int, taskCount)
		for {
			pa := <-pach // block by receive parameters
			go func() {
				thr <- 1 // occupy 1 thread
				piovit, err := partitionArray(ary, pa.left, pa.right)
				if err != nil {
					errch <- err
				}
				resultch <- &dataWrap{left: pa.left, right: pa.right, pivot: piovit}
				<-thr // release
			}()

		}
	}()
	size, err := ary.Len()
	if err != nil {
		return err
	}
	pach <- &param{left: 0, right: size}
	for {
		select {
		case err := <-errch:
			return err
		case re := <-resultch:
			if re.left < re.pivot {
				pach <- &param{left: re.left, right: re.pivot}
			}
			if re.pivot < re.right {
				pach <- &param{left: re.pivot, right: re.right}
			}
		}
	}

}
func MultTaskSort(ary Array, taskCount int) error {
	pch := make()
	return nil
}
