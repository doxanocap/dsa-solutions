package tiktok

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CardPackets(t *testing.T) {
	type testCase struct {
		cardTypes []int
		answer    int
	}

	tcs := []testCase{
		{
			cardTypes: []int{4, 7, 5, 11, 15},
			answer:    4,
		},
		{
			cardTypes: []int{3, 8, 7, 6, 4},
			answer:    2,
		},
	}

	for _, tc := range tcs {
		result := cardPackets(tc.cardTypes)
		if result == tc.answer {
			continue
		}
		assert.Equal(t, tc.answer, result)
	}
}

func Test_cardPackets(t *testing.T) {
	tests := []struct {
		cardTypes []int
		want      int
	}{
		{
			cardTypes: []int{4, 7, 5, 11, 15},
			want:      4,
		},
		{
			cardTypes: []int{3, 8, 7, 6, 4},
			want:      2,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("cardPackets-%d", i+1), func(t *testing.T) {
			assert.Equalf(t, tt.want, cardPackets(tt.args.cardTypes), "cardPackets(%v)", tt.cardTypes)
		})
	}
}
