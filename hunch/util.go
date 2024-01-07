package hunch

import (
	"context"
	"sort"
	"sync"
)

var (
	once               sync.Once
	poolerList         sync.Pool
	poolerIndexedValue sync.Pool
)

func init() {
	once.Do(func() {
		poolerList.New = func() interface{} {
			return make([]any, 0)
		}

		poolerIndexedValue.New = func() interface{} {
			return IndexedValue[any]{}
		}

	})
}

func pluckVals[T any](iVals []IndexedValue[T]) []T {
	//vals := make([]T, 0, len(iVals))

	vals := poolerList.Get().([]T)
	defer poolerList.Put(vals)

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

			val, err := exec(ctx)
			if err != nil {
				temp.Err = err
				output <- temp
				return
			}
			temp.Value = IndexedValue[T]{i, val}
			output <- temp
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
