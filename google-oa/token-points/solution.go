package main

func solution(points []int, tokens string) int {
	ans := 0
	n := len(points)

	for i := 0; i < n; i++ {
		if tokens[i] == 'T' {
			ans += points[i]
			if i > 0 && tokens[i-1] == 'T' {
				ans++
			}
		}
	}

	return ans
}
