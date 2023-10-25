package whatsapp

import "strings"

func RemoveItalicBold(msg string) (res string) {
	res = msg
	res = strings.ReplaceAll(res, "*", "")
	res = strings.ReplaceAll(res, "_", "")
	return
}
