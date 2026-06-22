package leetcode_57

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)          // количество существующих интервалов
	ans := make([][]int, 0, n+1) // создаем новый слайс для результата с cap=n+1

	idx := 0
	// добавляем все интервалы, которые находятся до newInterval (не пересекаются)
	for ; idx < n && intervals[idx][1] < newInterval[0]; idx++ {
		ans = append(ans, intervals[idx]) // добавляем в результат интервалы, которые заканчиваются до начала newInterval
	}

	// объединяем newInterval с пересекающимися интервалами
	for ; idx < n && intervals[idx][0] <= newInterval[1]; idx++ {
		newInterval[0] = min(newInterval[0], intervals[idx][0]) // обновляем начало нового интервала до минимального значения
		newInterval[1] = max(newInterval[1], intervals[idx][1]) // обновляем конец нового интервала до максимального значения
	}

	// добавляем объединенный (или обновленный) интервал в результат
	ans = append(ans, newInterval)
	// добавляем оставшиеся интервалы (после пересекающихся)
	return append(ans, intervals[idx:]...)
}
