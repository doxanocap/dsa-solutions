package leetcode_371

import (
	"fmt"
)

func getSum(a int, b int) int {
	arr := make([]int, 0, 1024)

	for a != 0 || b != 0 {
		if a&b == 1 {
			arr = append(arr, 2)
		} else if a|b == 0 {
			arr = append(arr, 1)
		} else {
			arr = append(arr, 0)
		}

		fmt.Println("a", a)
		fmt.Println("b", b)
		fmt.Println(a | b)
		a = a >> 1
		b = b >> 1
	}

	fmt.Println(arr)

	return 0
}
