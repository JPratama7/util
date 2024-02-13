package hunch

import (
	"context"
	"errors"
	"sync"
)

func runner[T any](ctx context.Context, idx int, take bool, mut *sync.Mutex, wg *sync.WaitGroup, exec Executable[T], done chan<- int, data *ExecutableOutput[T]) {
	defer func() {
		// TODO: IMPROVE THIS
		defer func() {
			recover()
		}()

		wg.Done()
		mut.Unlock()
		done <- idx
		close(done)
	}()

	v, err := exec(ctx)

	mut.Lock()
	data.Err = err
	if !take {
		return
	}
	data.Value = v
}

func run[T any](ctx context.Context, num int, execs ...Executable[T]) (val []T, err error) {

	wg := new(sync.WaitGroup)
	mut := new(sync.Mutex)

	fullres := make([]ExecutableOutput[T], len(execs))
	earlyDone := make(chan int)
	wgCh := make(chan *struct{})

	localCfg := copyGlobalCfg()

	for i, exec := range execs {
		wg.Add(1)
		go runner(ctx, i, num != 0, mut, wg, exec, earlyDone, &fullres[i])
	}

	defer func() {
		earlyDone = nil
		fullres = nil
	}()

	go func(wg *sync.WaitGroup, wgCh chan<- *struct{}) {
		wg.Wait()
		wgCh <- nil
		close(wgCh)
	}(wg, wgCh)

	if localCfg.forgetAll {
		return
	}

	if localCfg.earlyDone {
		select {
		case <-ctx.Done():
			err = errors.New("context canceled")
			return
		case idx := <-earlyDone:
			if fullres[idx].Err != nil {
				err = fullres[idx].Err
				return
			}

			val = make([]T, 0, 1)
			val = append(val, fullres[idx].Value)
			return
		}
	}

BREAKER:
	select {
	case <-ctx.Done():
		err = errors.New("context canceled")
		return
	case <-wgCh:
		break BREAKER
	}

	val, err = takeUntilEnough(num, num != 0, &localCfg, fullres...)
	return
}

func takeUntilEnough[T any](total int, take bool, cfg *globalConfig, res ...ExecutableOutput[T]) (uVals []T, err error) {
	totalData := len(res)

	if total != 0 {
		totalData = total
	}

	if take {
		uVals = make([]T, 0, totalData)
	}

	for _, r := range res[:totalData] {

		if r.Err != nil {
			if cfg.ignoreErr {
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
