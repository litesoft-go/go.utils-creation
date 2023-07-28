package go_utils_creation

import (
	"strconv"
	"testing"
)

func TestLinesToString(t *testing.T) {
	tests := []struct {
		expectedLines []string
		wantExpected  string
	}{
		{nil, ""},
		{[]string{}, ""},
		{[]string{""}, ""},
		{[]string{" "}, " "},
		{[]string{"", ""}, "\n"},
		{[]string{"- ", " +"}, "- \n +"},
		{[]string{"", "-", ""}, "-"},
		{[]string{"", "-", "+", "-", "+", ""}, "-\n+\n-\n+"},
	}
	for i, tt := range tests {
		name := strconv.Itoa(i)
		t.Run(name, func(t *testing.T) {
			if gotExpected := LinesToString(tt.expectedLines...); gotExpected != tt.wantExpected {
				t.Errorf("LinesToString() = %v, want %v", gotExpected, tt.wantExpected)
			}
		})
	}
}
