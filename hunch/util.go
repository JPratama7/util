package hunch

import (
	"context"
	"sort"
	"sync"
)

func pluckVals[T any](iVals []IndexedValue) []T {
	vals := make([]T, 0, len(iVals))

	for _, val := range iVals {
		conv, ok := val.Value.(T)
		if !ok {
			panic("hunch: type assertion failed")
		}

		vals = append(vals, conv)
	}

	return vals
}

func sortIdxVals(iVals []IndexedValue) (sorted []IndexedValue) {
	sorted = iVals

	sort.SliceStable(
		sorted,
		func(i, j int) bool {
			return sorted[i].Index < sorted[j].Index
		},
	)

	return sorted
}

func runExecs[T any](ctx context.Context, output chan<- IndexedExecutableOutput, execs []Executable[T]) {
	var wg sync.WaitGroup
	for i, exec := range execs {
		wg.Add(1)

		go func(i int, exec Executable[T]) {
			defer wg.Done()

			temp := IndexedExecutableOutput{}

			defer func() {
				output <- temp
			}()

			val, err := exec(ctx)
			if err != nil {
				temp.Err = err
				return
			}

			temp.Value = IndexedValue{i, val}
		}(i, exec)
	}

	wg.Wait()
	close(output)
}

func takeUntilEnough(fail chan<- error, success chan<- []IndexedValue, num int, output <-chan IndexedExecutableOutput, done chan<- bool, takeVal bool) {
	uVals := make([]IndexedValue, 0, num)

	enough := false
	outputCount := 0
	for r := range output {
		if enough {
			break
		}

		if r.Err != nil {
			enough = true
			fail <- r.Err
			break
		}

		if !takeVal {
			continue
		}

		uVals = append(uVals, r.Value)
		outputCount++

		if outputCount == num {
			enough = true
			success <- uVals
			break
		}
	}
	done <- true
}

func initVal[T any]() T {
	return *new(T)
}
