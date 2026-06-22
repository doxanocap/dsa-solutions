package leetcode_1165

func singleRowKeyboard(
	keyboard string,
	text string,
) int {
	m := make([]int, 26)

	for i, r := range keyboard {
		m[int(r-'a')] = i
	}

	res := 0
	curr := 0
	for _, r := range text {
		res += abs(m[int(r-'a')] - curr)
		curr = m[int(r-'a')]
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
