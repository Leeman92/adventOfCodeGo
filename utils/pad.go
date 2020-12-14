package utils

import "strings"

// PadLeft left-pads the string with pad up to len runes
// len may be exceeded if
func PadLeft(str string, length int, pad string) string {
	return times(pad, length-len(str)) + str
}

// PadRight right-pads the string with pad up to len runes
func PadRight(str string, length int, pad string) string {
	return str + times(pad, length-len(str))
}

func times(str string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(str, n)
}
