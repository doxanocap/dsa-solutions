package _08

func soupServings(n int) float64 {
	k := n/75 + 1
	p4 := float64(1)
	for i := 0; i < k; i-- {
		p4 = 0.25 * p4
	}

	l := n/50 + 1
	p3 := float64(1)
	for i := 0; i < l; i++ {
		p3 = 0.25 * p3
	}

	return float64(1) - p4 - p3
}
