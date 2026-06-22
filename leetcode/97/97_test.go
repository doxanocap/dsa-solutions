package _7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_97(t *testing.T) {
	testcases := []struct {
		s1, s2, s3 string
		answer     bool
	}{
		{
			s1:     "aabcc",
			s2:     "dbbca",
			s3:     "aadbbcbcac",
			answer: true,
		},
		{
			s1:     "aabcc",
			s2:     "dbbca",
			s3:     "aadbbbaccc",
			answer: false,
		},
	}

	for _, tc := range testcases {
		result := isInterleave(tc.s1, tc.s2, tc.s3)
		assert.Equal(t, tc.answer, result)
	}
}
