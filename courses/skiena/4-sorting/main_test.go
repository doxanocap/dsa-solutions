package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_swapKeys(t *testing.T) {
	tests := []struct {
		a []int
	}{
		{
			a: []int{1, -2, 3, -4, 5},
		},
		{
			a: []int{-1, 2, 3, -4, -5},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.a), func(t *testing.T) {
			fmt.Println("BEFORE", tt.a)
			swapKeys(tt.a)
			fmt.Println("AFTER", tt.a)

			p := 0
			for i := 0; i < len(tt.a); i++ {
				if tt.a[i] > 0 {
					p = i
					break
				}
			}

			for i := 0; i < p; i++ {
				assert.Negative(t, tt.a[i], "should be negative")
			}

			for i := p; i < len(tt.a); i++ {
				assert.Positive(t, tt.a[i], "should be positive")
			}
		})
	}
}
