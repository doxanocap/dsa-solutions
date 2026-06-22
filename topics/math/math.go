package math

// C
// combinations
func C(n, k int) int64 {
	if k < 0 || k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}

	if k > n/2 {
		k = n - k
	}

	result := int64(1)
	for i := 1; i <= k; i++ {
		result = result * int64(n-i+1) / int64(i)
	}

	return result
}

func GCD(a, b int64) int64 {
	for b != 0 {
		a %= b
		a, b = b, a
	}
	return a
}

func IsPrime(n int) bool {
	if n > 1 && n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func Pow(a, n int64) int64 {
	var res int64 = 1
	for n > 0 {
		if n&1 == 1 {
			res *= a
		}
		a *= a
		n >>= 1
	}
	return res
}
