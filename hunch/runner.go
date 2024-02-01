package hunch

import (
	"context"
	"sync"
)

func runner[T any](ctx context.Context, i int, take bool, mut *sync.Mutex, wg *sync.WaitGroup, exec Executable[T], data *IndexedExecutableOutput[T]) {
	mut.Lock()

	defer func() {
		wg.Done()
		mut.Unlock()
	}()

	val, err := exec(ctx)
	data.Err = err

	if !take {
		return
	}

	data.Value = IndexedValue[T]{i, val}
}

func run[T any](ctx context.Context, ignoreErr bool, num int, execs ...Executable[T]) (val []IndexedValue[T], err error) {

	wg := new(sync.WaitGroup)
	mut := new(sync.Mutex)

	fullres := make([]IndexedExecutableOutput[T], len(execs))

	for i, exec := range execs {
		wg.Add(1)
		go runner(ctx, i, num != 0, mut, wg, exec, &fullres[i])
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

	for _, r := range res[:totalData] {

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
