package sync

func defaultAlloc[T any]() any {
	return *new(T)
}
