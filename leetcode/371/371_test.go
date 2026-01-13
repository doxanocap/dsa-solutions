package leetcode_371

import (
	"fmt"
	"testing"
)

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				2, 3,
			},
			want: 3,
		},
		// TODO: Add test cases.
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("getSum-%d", i), func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
