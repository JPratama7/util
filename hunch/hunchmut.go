// Package hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive.
package hunch

import (
	"context"
)

// TakeMut Take returns the first `num` values outputted by the Executables.
func TakeMut[T any](parentCtx context.Context, num int, execs ...Executable[T]) ([]T, error) {
	execCount := len(execs)

	if num > execCount {
		num = execCount
	}

	// Create a new sub-context for possible cancelation.
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	i, err := run(ctx, num, execs...)
	if err != nil {
		return nil, err
	}

	return pluckVals(i), nil
}

// AllMut All returns all the outputs from all Executables, order guaranteed.
func AllMut[T any](parentCtx context.Context, execs ...Executable[T]) ([]T, error) {
	// Create a new sub-context for possible cancelation.
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	i, err := run(ctx, len(execs), execs...)
	if err != nil {
		return nil, err
	}

	return pluckVals(sortIdxVals(i)), nil
}

// ThrowMut Throw execute and ignore all the outputs from all Executables.
func ThrowMut[T any](parentCtx context.Context, execs ...Executable[T]) error {
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	_, err := run(ctx, 0, execs...)
	if err != nil {
		return err
	}

	return nil
}
