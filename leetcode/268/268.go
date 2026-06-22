package leetcode_268

func missingNumber(nums []int) int {
	// В Go a^b = a XOR b
	//
	// Truth table:
	// a | b | a ^ b
	// 0 | 1 |   1
	// 1 | 0 |   1
	// 1 | 1 |   0
	// 0 | 0 |   0
	// Для решений задачи мы должны использовать два свойства XOR операций:
	// 1. Последовательность XOR операций не имеет значения:
	//    (a^b)^c = a^(b^c)
	// 2. self-xor:
	//    x ^ x = 0
	// Возьмём XOR чисел [0,n] с числами из nums и получим недостающее число.
	// Наглядный пример:
	//    0 ^ 0 ^ 1 ^ 1 ^ 2 ^ 3 ^ 3 = 0 ^ 0 ^ 2 ^ 0 = 2
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		n = n ^ (nums[i] ^ i)
	}
	return n
}
