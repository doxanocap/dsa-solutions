package problems

func minDaysBloom(roses []int, k int, n int) int {
	if len(roses) < k*n {
		return -1
	}

	l, h := 1, 0
	for _, day := range roses {
		if day > h {
			h = day
		}
	}

	for l < h {
		mid := l + (h-l)/2

		c := 0
		bouquets := 0
		for _, day := range roses {
			if day <= mid {
				c++
				if c == mid {
					bouquets++
					c = 0
				}
			} else {
				c = 0
			}
		}

		if c == n {
			h = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
