package go_utils_creation

import (
	"fmt"
	"log"
	"testing"
)

func TestAsciiCharsN(t *testing.T) {
	tests := []struct {
		c    byte
		n    int
		want string
	}{
		{'c', 3, "ccc"},
		{'-', 1, "-"},
		{'|', 16, "||||||||||||||||"},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("'%v'x%v", tt.c, tt.n)
		t.Run(name, func(t *testing.T) {
			if got := AsciiCharsN(tt.c, tt.n); got != tt.want {
				t.Errorf("AsciiCharsN() = %v, want %v", got, tt.want)
			}
		})
	}
	expectPanic(t, 255, 1)
	expectPanic(t, '0', 0)
	expectPanic(t, '1', -1)
}

func expectPanic(t *testing.T, c byte, n int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("expected panic occurred:", err)
		}
	}()
	t.Errorf("AsciiCharsN() = %v, expected Panic", AsciiCharsN(c, n))
}
