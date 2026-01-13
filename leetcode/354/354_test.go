package _54

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_354(t *testing.T) {
	testcases := []struct {
		e      [][]int
		answer int
	}{
		{
			e: [][]int{
				{5, 4},
				{6, 4},
				{6, 7},
				{2, 3},
			},
			answer: 3,
		},
		{
			e: [][]int{
				{10, 8},
				{1, 12},
				{6, 15},
				{2, 18},
			},
			answer: 2,
		},
	}

	for _, tc := range testcases {
		result := maxEnvelopes(tc.e)
		assert.Equal(t, tc.answer, result)
	}
}
