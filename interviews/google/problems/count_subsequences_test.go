package problems

import (
	"fmt"
	"testing"
)

func Test_countSubsequences(t *testing.T) {
	tests := []struct {
		input []int
		want  int64
	}{
		{
			input: []int{1, 2, 3},
			want:  5,
		},
		{
			input: []int{4, 9, 3, 6, 5, 9},
			want:  12,
		},
	}
	for _, tt := range tests {

		t.Run(fmt.Sprintf("%v", tt.input),
			func(t *testing.T) {
				if got := countSubsequences(tt.input); got != tt.want {
					t.Errorf("countSubsequences() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
