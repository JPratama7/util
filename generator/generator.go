package generator

type Next interface {
	Next() bool
}

type HasNext interface {
	HasNext() bool
}

type Value[K comparable, T any] interface {
	Value() (K, T)
}

type Generator[K comparable, T any] interface {
	Next
	HasNext
	Value[K, T]
}
