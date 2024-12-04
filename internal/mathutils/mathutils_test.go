package mathutils

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{name: "positive number", input: 7, want: 7},
		{name: "negative number", input: -5, want: 5},
		{name: "zero", input: 0, want: 0},
		{name: "minus zero", input: -0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Abs(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
