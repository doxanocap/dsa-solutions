package uber

func solve2(n int) []int {
	if n == 1 {
		return []int{1}
	}

	prev := solve2(n / 2)

	newAns := make([]int, 0, n)
	for _, s := range prev {
		newAns = append(newAns, s, n+1-s)
	}

	return newAns
}

func solve3(n int) []int {
	if n == 1 {
		return []int{1}
	}

	prev := solve2(n / 2)

	newAns := make([]int, 0, n)
	for i, s := range prev {
		newAns = append(newAns, s, n/2+1+i)
	}

	return newAns
}
