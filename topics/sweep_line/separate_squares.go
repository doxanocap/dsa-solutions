package sweep_line

import "slices"

type Event struct {
	y, l, t int
}

func separateSquares(A [][]int) float64 {
	n := len(A)
	total := float64(0)
	events := make([]Event, 0, 2*n)

	for _, v := range A {
		events = append(events,
			Event{
				v[0], v[2], 0,
			},
			Event{
				v[0] + v[2], v[2], 1,
			})
		total += float64(v[2] * v[2])
	}

	slices.SortFunc(events, func(a, b Event) int {
		return a.y - b.y
	})

	area := float64(0)
	prevH := 0
	width := 0

	for _, event := range events {
		h := event.y
		temp := float64((h - prevH) * width)

		if area+temp >= total/2 {
			return float64(prevH) + (total/2-area)/float64(width)
		}

		area += temp
		if event.t == 0 {
			width += event.l
		} else {
			width -= event.l
		}
		prevH = h
	}
	return 0
}
