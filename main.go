package main

import (
	"fmt"
	"runtime"
	//"net/http/pprof"
)

// countries = [{UK, coins: 100}, {US, coins: 300}, ]
// 10 100 200

func main() {

	//if 1  {
	//	fmt.Println("1")
	//}

	coins := []int{1, 2, 3, 5}
	sum := 5

	res := get(coins, sum)

	fmt.Println(res)

	fmt.Println(runtime.NumGoroutine())
}

type C struct {
	A int
}

func (c C) String() string {
	return fmt.Sprintf("%d", c.A)
}

type Interface struct {
	Pointer *string
	Data    *string
}

func get(coins []int, sum int) int {
	// 1$,2$,3$,5$
	// [[2,3], [5], [1,3], [1,2]]

	count := 0

	var dfs func(idx, prev int)
	dfs = func(idx, curr int) {
		if curr > sum {
			return
		}
		if idx > len(coins)-1 {
			return
		}

		canExtend := false
		for i := idx; i < len(coins); i++ {
			if curr+coins[i] <= sum {
				canExtend = true
				dfs(i+1, curr+coins[i])
			}
		}

		if curr > 0 && !canExtend {
			count++
		}
	}

	dfs(0, 0)

	return count
}
