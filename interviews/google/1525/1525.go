package _525

func numSplits(s string) int {
	n := len(s)

	lm := map[byte]bool{}
	rm := map[byte]bool{}

	ld := make([]int, n)
	rd := make([]int, n)

	for i := 0; i < n; i++ {
		lm[s[i]] = true
		rm[s[n-1-i]] = true
		ld[i] = len(lm)
		rd[n-1-i] = len(rm)
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		if ld[i] == rd[i+1] {
			ans++
		}
	}
	return ans
}
