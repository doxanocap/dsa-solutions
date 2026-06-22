package amazon

// P1
func minDeliveryTime(n int, orderCityList []int) int {
	// Подсчёт заказов на каждый город
	count := make([]int, n+1)
	for _, city := range orderCityList {
		count[city]++
	}

	// Бинарный поиск по минимальному времени
	lo, hi := 1, len(orderCityList)*2
	for lo < hi {
		mid := (lo + hi) / 2
		if canFinish(count, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// Проверка: можно ли выполнить все заказы за T дней
func canFinish(count []int, T int) bool {
	n := len(count) - 1
	extra := 0 // сколько "лишних" заказов нужно отправить в другие центры

	for i := 1; i <= n; i++ {
		if count[i] > T {
			// лишние заказы
			extra += count[i] - T
		}
	}

	// Сколько свободных "окон" есть у других центров (по 2 дня за заказ)
	spare := 0
	for i := 1; i <= n; i++ {
		if count[i] < T {
			spare += (T - count[i]) / 2
		}
	}

	return spare >= extra
}

// P2
// Возвращает кратчайшую (и при равной длине лексикографически минимальную)
// строку цифр, переводящую цифру start в цифру goal за шаги +x или +y (mod 10).
// Строка НЕ содержит стартовую цифру, только последующие цифры каждого шага.
// ok=false, если достигнуть goal нельзя.
func shortestPath(start, goal, x10, y10 int) (path string, ok bool) {
	const INF = int(1e9)
	// dist[d] — минимальное число шагов до цифры d
	dist := make([]int, 10)
	for i := range dist {
		dist[i] = INF
	}
	// best[d] — лексикографически минимальная строка цифр для dist[d]
	best := make([]string, 10)

	type node struct{ d int }
	q := []node{{start}}
	dist[start] = 0
	best[start] = ""

	head := 0
	push := func(v int, candStr string, candDist int) {
		if candDist < dist[v] || (candDist == dist[v] && candStr < best[v]) {
			dist[v] = candDist
			best[v] = candStr
			q = append(q, node{v})
		}
	}

	for head < len(q) {
		u := q[head].d
		head++

		// два возможных перехода
		v1 := (u + x10) % 10
		s1 := best[u] + string('0'+v1)
		push(v1, s1, dist[u]+1)

		v2 := (u + y10) % 10
		s2 := best[u] + string('0'+v2)
		push(v2, s2, dist[u]+1)
	}

	if dist[goal] == INF {
		return "", false
	}
	return best[goal], true
}

func missingDigits(config string, x, y int32) string {
	if len(config) == 0 {
		// По условию, если нельзя восстановить — вернуть "1".
		// Для пустого ввода разумно вернуть "1".
		return "1"
	}
	// Работают только остатки по модулю 10.
	x10 := int((x%10 + 10) % 10)
	y10 := int((y%10 + 10) % 10)

	// Результат начинается с первой наблюдаемой цифры.
	res := string(config[0])

	for i := 0; i+1 < len(config); i++ {
		start := int(config[i] - '0')
		goal := int(config[i+1] - '0')

		path, ok := shortestPath(start, goal, x10, y10)
		if !ok {
			return "1" // по условию — если нельзя восстановить
		}
		// path уже содержит последовательность цифр шагов (без стартовой цифры),
		// заканчивается goal. Просто дописываем.
		res += path
	}
	return res
}
