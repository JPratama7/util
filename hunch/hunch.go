// Package hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive.
package hunch

import (
	"context"
	"fmt"
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

	output := make(chan IndexedExecutableOutput[T], 1)
	go runExecs(ctx, output, execs)

	fail := make(chan error, 1)
	done := make(chan bool)
	success := make(chan []IndexedValue[T], 1)
	go takeUntilEnough(fail, success, min(len(execs), num), output, done, true)

	select {

	case <-parentCtx.Done():
		// Stub comment to fix a test coverage bug.
		return nil, parentCtx.Err()

	case err := <-fail:
		cancel()
		if parentCtxErr := parentCtx.Err(); parentCtxErr != nil {
			return nil, parentCtxErr
		}
		return nil, err

	case uVals := <-success:
		cancel()
		return pluckVals[T](uVals), nil

	case <-done:
		cancel()
		return nil, nil
	}
}

// All returns all the outputs from all Executables, order guaranteed.
func All[T any](parentCtx context.Context, execs ...Executable[T]) ([]T, error) {
	// Create a new sub-context for possible cancelation.
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	output := make(chan IndexedExecutableOutput[T], 1)
	fail := make(chan error, 1)
	success := make(chan []IndexedValue[T], 1)
	done := make(chan bool)

	go runExecs(ctx, output, execs)
	go takeUntilEnough(fail, success, len(execs), output, done, true)

	select {

	case <-parentCtx.Done():
		// Stub comment to fix a test coverage bug.
		return nil, parentCtx.Err()

	case err := <-fail:
		if parentCtxErr := parentCtx.Err(); parentCtxErr != nil {
			return nil, parentCtxErr
		}
		return nil, err

	case uVals := <-success:
		return pluckVals(sortIdxVals(uVals)), nil

	case <-done:
		cancel()
		return nil, nil
	}
}

// Throw execute and ignore all the outputs from all Executables.
func Throw[T any](parentCtx context.Context, execs ...Executable[T]) error {
	// Create a new sub-context for possible cancelation.
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	output := make(chan IndexedExecutableOutput[T], 1)
	fail := make(chan error, 1)
	success := make(chan []IndexedValue[T], 1)
	done := make(chan bool)

	go runExecs(ctx, output, execs)
	go takeUntilEnough(fail, success, len(execs), output, done, false)

	select {

	case <-parentCtx.Done():
		// Stub comment to fix a test coverage bug.
		return parentCtx.Err()

	case err := <-fail:
		if parentCtxErr := parentCtx.Err(); parentCtxErr != nil {
			return parentCtxErr
		}
		return err

	case <-done:
		cancel()
		return nil
	}
}

/*
Last returns the last `num` values outputted by the Executables.
*/
func Last[T any](parentCtx context.Context, num int, execs ...Executable[T]) ([]T, error) {
	execCount := len(execs)
	if num > execCount {
		num = execCount
	}
	start := execCount - num

	vals, err := Take(parentCtx, execCount, execs...)
	if err != nil {
		return nil, err
	}

	return vals[start:], err
}

// MaxRetriesExceededError stores how many times did an Execution run before exceeding the limit.
// The retries field holds the value.
type MaxRetriesExceededError struct {
	retries int
}

func (err MaxRetriesExceededError) Error() string {
	var word string
	switch err.retries {
	case 0:
		word = "infinity"
	case 1:
		word = "1 time"
	default:
		word = fmt.Sprintf("%v times", err.retries)
	}

	return fmt.Sprintf("Max retries exceeded (%v).\n", word)
}

// Retry attempts to get a value from an Executable instead of an Error.
// It will keep re-running the Executable when failed no more than `retries` times.
// Also, when the parent Context canceled, it returns the `Err()` of it immediately.
func Retry[T any](parentCtx context.Context, retries int, fn Executable[T]) (T, error) {
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	c := 0
	fail := make(chan error, 1)
	success := make(chan T, 1)

	for {
		go func() {
			val, err := fn(ctx)
			if err != nil {
				fail <- err
				return
			}
			success <- val
		}()

		select {
		//
		case <-parentCtx.Done():
			// Stub comment to fix a test coverage bug.
			return initVal[T](), parentCtx.Err()

		case <-fail:
			if parentCtxErr := parentCtx.Err(); parentCtxErr != nil {
				return initVal[T](), parentCtxErr
			}

			c++
			if retries == 0 || c < retries {
				continue
			}
			return initVal[T](), MaxRetriesExceededError{c}

		case val := <-success:
			return val, nil
		}
	}
}

// Waterfall runs `ExecutableInSequence`s one by one,
// passing previous result to next Executable as input.
// When an error occurred, it stop the process then returns the error.
// When the parent Context canceled, it returns the `Err()` of it immediately.
func Waterfall[T any](parentCtx context.Context, execs ...ExecutableInSequence[T]) (T, error) {
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	var lastVal T
	execCount := len(execs)
	i := 0
	fail := make(chan error, 1)
	success := make(chan T, 1)

	for {
		go func() {
			val, err := execs[i](ctx, lastVal)
			if err != nil {
				fail <- err
				return
			}
			success <- val
		}()

		select {

		case <-parentCtx.Done():
			// Stub comment to fix a test coverage bug.
			return initVal[T](), parentCtx.Err()

		case err := <-fail:
			if parentCtxErr := parentCtx.Err(); parentCtxErr != nil {
				return initVal[T](), parentCtxErr
			}

			return initVal[T](), err

		case val := <-success:
			lastVal = val
			i++
			if i == execCount {
				return val, nil
			}

			continue
		}
	}
}
