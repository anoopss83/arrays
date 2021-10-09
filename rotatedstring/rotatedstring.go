package rotatedstring

import "strings"

// isRotated returns true if s2 is a rotated version of s1
func isRotated(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s2 = s2 + s2
	return strings.Contains(s2, s1)
}
