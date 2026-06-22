package mocks

import (
	"fmt"
	"math"
	"slices"
	"strconv"
)

func completePrime(num int) bool {
	A := make([]int, 0)

	x := num
	for x > 0 {
		d := x % 10

		A = append(A, d)
		x = x / 10
	}

	for i := 0; i < len(A); i++ {

		l := toNum(A[:i])
		r := toNum(A[i+1:])

		if !isPrime(l) || !isPrime(r) {
			return false
		}
	}

	return true
}

func toNum(s []int) int {
	n := len(s)

	x := int(0)
	for i := n - 1; i >= 0; i-- {
		x = x*10 + s[i]
	}

	return x
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}

	f := int(math.Sqrt(float64(n)))

	for i := 2; i <= f; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func binary(x int) []int {

	bs := make([]int, 0)
	for x > 0 {
		b := x & 1

		bs = append(bs, b)
		x >>= 1
	}

	slices.Reverse(bs)
	return bs
}

func minOperations(nums []int) []int {
	//maxN := 5000
	//
	//pals := []int{}
	//for x := 1; i <= maxN; x++ {
	//	if isPalBinary(x) {
	//		pals = append(pals, x)
	//	}
	//}
	//
	//ans := make([]int, len(nums))
	//for i := range nums {
	//	j := sort.Searchints(pals, num)
	//	j = min(j, len(pals)-1)
	//
	//	diff := abs(num - pals[j])
	//	if j > 0 {
	//		diff = min(diff, abs(num-pals[j-1]))
	//	}
	//	ans[i] = diff
	//}
	//
	//sort.SearchInts()
	//
	//return ans
}

func reversBinary(n int) int {
	bs := make([]int, 0, n)
	for n > 0 {
		b := n & 1
		bs = append(bs, b)
		n >>= 1
	}

	j := 0
	for j < len(bs) && bs[j] == 0 {
		j++
	}

	x := 0
	pow2 := 1
	for i := len(bs) - 1; i >= j; i-- {
		if bs[i] == 1 {
			x += pow2
		}

		pow2 = pow2 * 2
	}
	return x
}

func isPalBinary(x int) bool {
	b := []byte(strconv.FormatInt(int64(x), 2))

	i, j := 0, len(b)-1
	for i < j {
		if b[i] != b[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxPoints(t1 []int, t2 []int, k int) int64 {
	n := len(t1)

	events := make([][2]int, n)

	for i := 0; i < n; i++ {
		diff := t1[i] - t2[i]

		events[i] = [2]int{diff, i}
	}

	slices.SortFunc(events, func(a, b [2]int) int {
		if a[0] == b[0] {
			i, j := a[1], b[1]
			return i - j
		}
		return a[0] - b[0]
	})

	fmt.Println(events)

	return 0
}
