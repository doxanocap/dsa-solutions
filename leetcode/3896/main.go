package _896

import (
	"math/big"
	"slices"
)

var mod = int64(1e9 + 7)

func maxValue(nums1 []int, nums0 []int) int {
	n := len(nums1)

	pairs := make([][2]int, n)
	for i := range pairs {
		pairs[i] = [2]int{nums1[i], nums0[i]}
	}

	slices.SortFunc(pairs, func(a, b [2]int) int {
		la := a[0] + a[1]
		lb := b[0] + b[1]

		powLb := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(lb)), nil)
		powLa := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(la)), nil)

		segA := new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(a[0])), nil), big.NewInt(1))
		segA.Lsh(segA, uint(a[1]))

		segB := new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(b[0])), nil), big.NewInt(1))
		segB.Lsh(segB, uint(b[1]))

		AB := new(big.Int).Add(new(big.Int).Mul(segA, powLb), segB)
		BA := new(big.Int).Add(new(big.Int).Mul(segB, powLa), segA)

		return BA.Cmp(AB)
	})

	ans := int64(0)

	for _, p := range pairs {
		ones := p[0]
		zeros := p[1]

		val := (pow(2, ones) - 1 + mod) % mod
		val = (val * pow(2, zeros)) % mod

		ln := ones + zeros

		ans = (ans * pow(2, ln)) % mod
		ans = (ans + val) % mod
	}

	return int(ans)
}

func pow(base int64, exp int) int64 {
	// Обработка отрицательной степени
	if exp < 0 {
		base = 1 / base
		exp = -exp
	}

	res := int64(1)
	for exp > 0 {
		// Если текущий бит степени равен 1 (степень нечётная)
		if exp%2 == 1 {
			res *= base
		}
		base *= base // Возводим основание в квадрат
		exp /= 2     // Сдвигаем степень вправо (делим на 2)
	}
	return res
}
