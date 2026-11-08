package main

import "testing"

func Test_part2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input string
		want  int
	}{
		{
			name:  "test",
			input: "2x3x4\n1x1x10",
			want:  48,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := part2(tt.input)
			if got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
