package main

import (
	"fmt"
	"testing"
)

func Test_solution(t *testing.T) {
	m := map[int]int{}

	for i := 0; i < 10_000; i++ {
		x1 := i / 1000
		x2 := (i / 100) % 10
		x3 := (i % 100) / 10
		x4 := i % 10
		m[x1+x2+x3+x4]++
	}

	fmt.Println(m)

	tests := make([]int, 36)
	for i := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			if got := solution(i); got != m[i] {
				t.Errorf("solution() = %v, want %v", got, m[i])
			}
		})
	}
}
