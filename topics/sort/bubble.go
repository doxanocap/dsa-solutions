package sort

func Bubble(A []int) {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A)-i-1; j++ {
			if A[j] > A[j+1] {
				A[j], A[j+1] = A[j+1], A[j]
			}
		}
	}
}
