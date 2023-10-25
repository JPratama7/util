package convert

func ToPtr[T any](b T) *T {
	return &b
}
