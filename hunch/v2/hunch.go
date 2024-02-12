// Package hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive.
package hunch

import (
	"context"
	"log"
)

// Take returns the first `num` values outputted by the Executables.
func Take[T any](parentCtx context.Context, num int, execs ...Executable[T]) ([]T, error) {
	execCount := len(execs)

	if num > execCount {
		num = execCount
	}

	// Create a new sub-context for possible cancelation.
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	resetGlobalCfg()

	i, err := run(ctx, num, execs...)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// All returns all the outputs from all Executables, order guaranteed.
func All[T any](parentCtx context.Context, execs ...Executable[T]) ([]T, error) {
	// Create a new sub-context for possible cancelation.
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	resetGlobalCfg()

	i, err := run(ctx, len(execs), execs...)
	log.Printf("i: %+v", err)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// Throw execute and ignore all the outputs from all Executables.
func Throw[T any](parentCtx context.Context, execs ...Executable[T]) {
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	resetGlobalCfg()
	SetForgetAll(true)

	_, _ = run(ctx, 0, execs...)

	return
}

// First execute and wait for first
func First[T any](parentCtx context.Context, execs ...Executable[T]) (val T, err error) {
	// Create a new sub-context for possible cancelation.
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	resetGlobalCfg()
	SetEarlyDone(true)

	i, err := run(ctx, 1, execs...)
	if err != nil {
		return
	}

	return i[0], nil

}
