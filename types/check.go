package types

import "reflect"

func IsPointer(b any) bool {
	if b == nil {
		return false
	}

	return reflect.TypeOf(b).Kind() == reflect.Ptr
}

func IsIterable(b any) bool {
	if b == nil {
		return false
	}
	return reflect.TypeOf(b).Kind() == reflect.Slice || reflect.TypeOf(b).Kind() == reflect.Array || reflect.TypeOf(b).Kind() == reflect.Map
}
