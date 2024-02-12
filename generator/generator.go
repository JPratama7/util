package generator

import "context"

type Next interface {
	Next() bool
}

type HasNext interface {
	HasNext() bool
}

type Value[K comparable, T any] interface {
	Value() (K, T)
}

type Chan[K comparable, T any] interface {
	Chan(context.Context) (chan K, chan T)
}

type Generator[K comparable, T any] interface {
	Next
	HasNext
	Value[K, T]
	Chan[K, T]
}
