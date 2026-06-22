package tiktok

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SecurityRisk(t *testing.T) {
	type testCase struct {
		videoHash     string
		encryptionKey string
		answer        int
	}

	tcs := []testCase{
		{
			videoHash:     "upload",
			encryptionKey: "doupla",
			answer:        1,
		},
	}

	for _, tc := range tcs {
		result := findSecurityRisk(tc.videoHash, tc.encryptionKey)
		assert.Equal(t, tc.answer, result)
	}
}
