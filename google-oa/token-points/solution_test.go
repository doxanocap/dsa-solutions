package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		points []int
		tokens string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.points, tt.args.tokens); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
