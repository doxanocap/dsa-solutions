package bits

import "math/bits"

func operations() {

	// mask | (1 << j) — Включение (Set bit)
	// mask & ~(1 << j) — Выключение (Clear bit)

	// (mask >> i) & 1 == 1 - сравнивание i-го бита

	bits.OnesCount(12) // считает 1 биты в маске
}
