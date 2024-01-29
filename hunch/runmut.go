package hunch

import (
	"context"
	"sync"
)

func runner[T any](ctx context.Context, i int, take bool, wg *sync.WaitGroup, mut *sync.RWMutex, exec Executable[T], res []IndexedExecutableOutput[T]) {
	mut.Lock()
	data := IndexedExecutableOutput[T]{}
	defer func() {
		res = append(res, data)
		mut.Unlock()
		wg.Done()
	}()

	val, err := exec(ctx)
	if err != nil {
		data.Err = err
		return
	}

	if !take {
		return
	}

	data.Value = IndexedValue[T]{i, val}
}

func run[T any](ctx context.Context, ignoreErr bool, num int, execs ...Executable[T]) (val []IndexedValue[T], err error) {

	mut := new(sync.RWMutex)
	wg := new(sync.WaitGroup)

	fullres := make([]IndexedExecutableOutput[T], 0, len(execs))

	for i, exec := range execs {
		wg.Add(1)
		go runner(ctx, i, num != 0, wg, mut, exec, fullres)
	}

	wg.Wait()

	val, err = takeUntilEnoughMut(num, num != 0, ignoreErr, fullres...)
	return
}

func takeUntilEnoughMut[T any](total int, take, ignoreErr bool, res ...IndexedExecutableOutput[T]) (uVals []IndexedValue[T], err error) {
	totalData := len(res)

	if total != 0 {
		totalData = total
	}

	if take {
		uVals = make([]IndexedValue[T], 0, totalData)
	}

	for _, r := range res {

		if r.Err != nil {
			if ignoreErr {
				continue
			}

			err = r.Err
			break
		}

		if take {
			uVals = append(uVals, r.Value)
		}
	}
	return
}
