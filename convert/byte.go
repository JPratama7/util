package convert

import "unsafe"

func UnsafeBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func SafeBytes(s string) []byte {
	b := UnsafeBytes(s)
	tmp := make([]byte, len(b))
	copy(tmp, b)
	return tmp
}
