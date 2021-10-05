package compressstring

import (
	"strconv"
	"strings"
)

func compress(s string) string {
	count := 1
	var out strings.Builder

	if countCompressedStr(s) >= len(s) {
		return s
	}

	i := 1
	for i < len(s) {
		if s[i-1] != s[i] {
			out.WriteByte(s[i-1])
			out.WriteString(strconv.Itoa(count))
			count = 1
		} else {
			count++
		}
		i++
	}
	out.WriteByte(s[i-1])
	out.WriteString(strconv.Itoa(count))
	return out.String()
}

func countCompressedStr(s string) int {
	n := 0
	i := 0
	for i+1 < len(s) {
		if s[i] != s[i+1] {
			n += 1
		}
		i++
	}
	return (n + 1) * 2
}
