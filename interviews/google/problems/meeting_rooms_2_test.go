package problems

import (
	"testing"
)

func Test_minChairs(t *testing.T) {
	type args struct {
		s []int
		e []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: []int{1, 2, 6, 5, 3},
				e: []int{5, 5, 7, 6, 8},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minChairs(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("minChairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
