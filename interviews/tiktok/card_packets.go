package tiktok

func cardPackets(cardTypes []int) int {
	minNum := cardTypes[0]
	for _, cardType := range cardTypes {
		minNum = min(minNum, cardType)
	}

	pns := primeNums(minNum)
	minSum := 99999999

	for _, pn := range pns {
		sum := 0
		for _, n := range cardTypes {
			if n%pn != 0 {
				sum += pn - n%pn
			}
		}
		minSum = min(minSum, sum)
	}

	return minSum
}

func primeNums(n int) []int {
	sieve := make([]bool, n+1)
	pns := []int{}

	for i := 2; i <= n; i++ {
		if !sieve[i] {
			pns = append(pns, i)
			for j := i * i; j <= n; j += i {
				sieve[j] = true
			}
		}
	}
	return pns
}
