package types

import "reflect"

func IsPointer(b any) bool {
	if b == nil {
		return false
	}

	return reflect.TypeOf(b).Kind() == reflect.Ptr
}
