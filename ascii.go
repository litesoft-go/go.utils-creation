package go_utils_creation

import "fmt"

// AsciiCharsN -- generates a 'string' made up of 'n' (single byte) characters, e.g. Ascii only!
// If 'c' is not ascii or 'n' < 1, a "panic" is thrown.
// -
func AsciiCharsN(c byte, n int) string {
	uc := uint8(c)
	if 127 < uc {
		panic(fmt.Sprintf("expected Ascii char/byte (0-127), but was: %v", uc))
	}
	if n < 1 {
		panic(fmt.Sprintf("at least 1 '%c' character must be requested, but request was for %v", c, n))
	}
	s := make([]byte, n)
	for i := 0; i < n; i++ {
		s[i] = c
	}
	return string(s)
}
