package main

import (
	"math/rand"
)

func main() {

}

func quickSelectArr(a []int, k int) []int {
	if len(a) <= k {
		return a
	}

	p := a[rand.Intn(len(a))]

	var (
		lows   []int
		pivots []int
		highs  []int
	)

	for _, v := range a {
		switch {
		case v == p:
			pivots = append(pivots, v)
		case v < p:
			lows = append(lows, v)
		case v > p:
			highs = append(highs, v)
		}
	}

	switch {
	case k <= len(lows):
		return quickSelectArr(lows, k)
	case k <= len(lows)+len(pivots):
		return append(lows, pivots[0:k-len(lows)]...)
	default:
		return append(append(lows, pivots...), quickSelectArr(highs, k-len(pivots)-len(lows))...)
	}
}

func quickSelect(a []int, k int) int {
	if len(a) == 1 {
		return a[0]
	}

	pivot := a[rand.Intn(len(a))]
	var lows, highs, pivots []int

	for _, v := range a {
		switch {
		case v < pivot:
			lows = append(lows, v)
		case v > pivot:
			highs = append(highs, v)
		default:
			pivots = append(pivots, v)
		}
	}

	switch {
	case k < len(lows):
		return quickSelect(lows, k)
	case k < len(lows)+len(pivots):
		return pivot
	default:
		return quickSelect(highs, k-len(lows)-len(pivots))
	}
}

func findMedian(a []int) float64 {
	n := len(a)
	if n%2 == 1 {
		return float64(quickSelect(a, n/2))
	}
	left := quickSelect(a, n/2-1)
	right := quickSelect(a, n/2)
	return float64(left+right) / 2.0
}

func mergeInplace(a, b []int) {
	l, r := 0, len(b)-1
	for l < len(a) && a[l] < 0 {
		l++
	}
	if l == len(a) {
		return
	}

	for r >= 0 && b[r] > 0 {
		r--
	}
	if r < 0 {
		return
	}

	for l < len(a) && r >= 0 {
		a[l], b[r] = b[r], a[l]
		r--
	}
}

func swapKeys(a []int) {
	if len(a) <= 1 {
		return
	} else if len(a) == 2 {
		if a[1] < a[0] {
			a[0], a[1] = a[1], a[0]
		}
		return
	}

	m := len(a) / 2

	left := a[0:m]
	swapKeys(left)

	right := a[m:]
	swapKeys(right)

	mergeInplace(left, right)
}

func swapKeysOn(a []int) {
	i := 0
	for j := 0; j < len(a); j++ {
		if a[j] < a[i] {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
}
