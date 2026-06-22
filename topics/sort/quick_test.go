package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func generateRandomArray(size int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(1000)
	}
	return arr
}

func TestQuickSortWithRandomInput(t *testing.T) {
	for i := 0; i < 10; i++ {
		size := rand.Intn(100)
		input := generateRandomArray(size)

		expected := make([]int, len(input))
		copy(expected, input)

		sort.Ints(expected)

		QuickSort(input)

		if !reflect.DeepEqual(input, expected) {
			t.Errorf("Test failed for input %v. QuickSort() = %v, expected %v", expected, input, expected)
		}
	}
}
