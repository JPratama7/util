package strings

import "github.com/JPratama7/util/convert"

func ToLower(b string) string {
	res := make([]byte, len(b))
	copy(res, b)
	for i := 0; i < len(res); i++ {
		res[i] = toLowerTable[res[i]]
	}

	return convert.UnsafeString(res)
}

// ToUpper converts ascii string to upper-case
func ToUpper(b string) string {
	res := make([]byte, len(b))
	copy(res, b)
	for i := 0; i < len(res); i++ {
		res[i] = toUpperTable[res[i]]
	}

	return convert.UnsafeString(res)
}
