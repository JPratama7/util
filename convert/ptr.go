package convert

func ToPointer[T any](b T) *T {
	return &b
}
