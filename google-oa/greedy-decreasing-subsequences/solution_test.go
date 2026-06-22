package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				A: []int{5, 4, 3, 6, 1},
			},
			want: 2,
		},
		{
			args: args{
				A: []int{10, 6, 9, 5, 8, 4, 7},
			},
			want: 2,
		},
		{
			args: args{
				A: []int{10, 11, 12, 14, 17, 9, 8, 1},
			},
			want: 5,
		},
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.A); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
