package hunch

import (
	"context"
)

// Executable represents a singular logic block.
// It can be used with several functions.
type Executable[T any] func(context.Context) (T, error)

// ExecutableInSequence represents one of a sequence of logic blocks.
type ExecutableInSequence[T any] func(context.Context, T) (T, error)

// ExecutableOutput stores both output and error values from a Excetable.
type ExecutableOutput[T any] struct {
	Value T
	Err   error
}

type globalConfig struct {
	forgetAll   bool
	ignoreErr   bool
	onlySuccess bool
	earlyDone   bool
	takeVal     bool
}
