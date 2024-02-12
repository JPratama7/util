package hunch

import (
	"context"
	"errors"
	"log"
	"sync"
)

func runner[T any](ctx context.Context, idx int, take bool, mut *sync.Mutex, wg *sync.WaitGroup, exec Executable[T], done chan<- int, data *ExecutableOutput[T]) {
	v, err := exec(ctx)

	defer func(a T, e error) {
		mut.Lock()
		data.Value = a
		data.Err = e
		mut.Unlock()
		wg.Done()
		done <- idx
	}(v, err)

	if !take {
		return
	}
}

func run[T any](ctx context.Context, num int, execs ...Executable[T]) (val []T, err error) {

	wg := new(sync.WaitGroup)
	mut := new(sync.Mutex)

	fullres := make([]ExecutableOutput[T], len(execs))
	earlyDone := make(chan int)
	wgCh := make(chan int, 1)

	localCfg := copyGlobalCfg()

	for i, exec := range execs {
		wg.Add(1)
		go runner(ctx, i, num != 0, mut, wg, exec, earlyDone, &fullres[i])
	}

	go func(wg *sync.WaitGroup, wgCh chan<- int) {
		wg.Wait()
		wgCh <- 1
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
			val = make([]T, 0, 1)
			val = append(val, fullres[idx].Value)
			log.Println("early done", val)
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

	val, err = takeUntilEnough(num, num != 0, localCfg, fullres...)
	return
}

func takeUntilEnough[T any](total int, take bool, cfg globalConfig, res ...ExecutableOutput[T]) (uVals []T, err error) {
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
