package types

func Sizer(s ...int) int {
	size := 0
	if len(s) > 0 {
		size = s[0]
	}
	return size
}
