package go_utils_creation

import "testing"

func Test_niy(t *testing.T) {
	tests := []struct {
		expected   string
		argMethod  string
		argComment string
		argParams  []any
	}{
		{"Test()", "Test", "", nil},
		{"Test(0) // note...", "Test", "note...", []any{0}},
		{"Testy('fred', 1)", "Testy", "", []any{"fred", 1}},
	}
	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := niy(tt.argMethod, tt.argComment, tt.argParams...); got != tt.expected {
				t.Errorf("niy() = %v, want %v", got, tt.expected)
			}
		})
	}
}
