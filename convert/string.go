package convert

import "unsafe"

func UnsafeString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func SafeString(b []byte) string {
	return string(b)
}
