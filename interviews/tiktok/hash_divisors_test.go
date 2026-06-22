package tiktok

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countHashtagDivisors(t *testing.T) {
	type args struct {
		description1 string
		description2 string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				description1: "ababab",
				description2: "ab",
			},
			want: 1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("hashDivisors-%d", i+1), func(t *testing.T) {
			assert.Equalf(t, tt.want, countHashtagDivisors(tt.args.description1, tt.args.description2), "countHashtagDivisors(%v, %v)", tt.args.description1, tt.args.description2)
		})
	}
}
