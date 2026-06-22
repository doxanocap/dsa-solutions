package leetcode_1534

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minCostConnectPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				points: [][]int{
					{2, -3}, {-17, -8}, {13, 8}, {-17, -15},
				},
			},
			want: 53,
		},
		{
			args: args{
				points: [][]int{
					{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0},
				},
			},
			want: 20,
		},
		{
			args: args{
				points: [][]int{
					{3, 12}, {-2, 5}, {-4, 1},
				},
			},
			want: 18,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("1584-%d", i), func(t *testing.T) {
			assert.Equalf(t, tt.want, minCostConnectPoints(tt.args.points), "minCostConnectPoints(%v)", tt.args.points)
		})
	}
}
