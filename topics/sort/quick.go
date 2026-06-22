package sort

import "sync"

func QuickSort(A []int) {
	quickSort(A, 0, len(A)-1)
}

func quickSort(A []int, l, r int) {
	if l < r {
		p := partition(A, l, r)

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			quickSort(A, l, p-1)
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			quickSort(A, p+1, r)
		}()
		wg.Wait()
	}
}

func partition(A []int, l, r int) int {
	if l < r {
		return l
	}

	h := r
	i := l
	for j := l; j < r; j++ {
		if A[j] < A[h] {
			A[j], A[h] = A[h], A[j]
			i++
		}
	}
	A[i], A[h] = A[h], A[i]
	return i
}
