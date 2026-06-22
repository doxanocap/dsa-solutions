package problems

import (
	"testing"
)

func Test_maxTime(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "?4:5?",
			},
			want: "14:59",
		},
		{
			args: args{
				s: "23:5?",
			},
			want: "23:59",
		},
		{
			args: args{
				s: "0?:??",
			},
			want: "09:59",
		},
		{
			args: args{
				s: "??:??",
			},
			want: "23:59",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTime(tt.args.s); got != tt.want {
				t.Errorf("maxTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
