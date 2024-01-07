package hunch

import (
	"context"
	"sort"
	"sync"
)

func pluckVals[T any](iVals []IndexedValue[T]) []T {
	vals := make([]T, 0, len(iVals))

	for _, val := range iVals {
		vals = append(vals, val.Value)
	}

	return vals
}

func sortIdxVals[T any](iVals []IndexedValue[T]) (sorted []IndexedValue[T]) {
	sorted = iVals

	sort.SliceStable(
		sorted,
		func(i, j int) bool {
			return sorted[i].Index < sorted[j].Index
		},
	)

	return sorted
}

func runExecs[T any](ctx context.Context, output chan<- IndexedExecutableOutput[T], execs []Executable[T]) {
	var wg sync.WaitGroup
	for i, exec := range execs {
		wg.Add(1)

		go func(i int, exec Executable[T]) {
			defer wg.Done()

			temp := IndexedExecutableOutput[T]{}

			defer func() {
				output <- temp
			}()

			val, err := exec(ctx)
			if err != nil {
				temp.Err = err
				return
			}

			temp.Value = IndexedValue[T]{i, val}
		}(i, exec)
	}

	wg.Wait()
	close(output)
}

func takeUntilEnough[T any](fail chan error, success chan []IndexedValue[T], num int, output chan IndexedExecutableOutput[T]) {
	uVals := make([]IndexedValue[T], 0, num)

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

		uVals = append(uVals, r.Value)
		outputCount++

		if outputCount == num {
			enough = true
			success <- uVals
			break
		}
	}
}

func initVal[T any]() T {
	return *new(T)
}
