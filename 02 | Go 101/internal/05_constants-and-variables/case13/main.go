package main

func main() {
	const (
		k = 3 // 在此处，iota == 0

		m float32 = iota + .5 // m float32 = 1 + .5
		n                     // n float32 = 2 + .5

		p = 9             // 在此处，iota == 3
		q = iota * 2      // q = 4 * 2
		_                 // _ = 5 * 2
		r                 // r = 6 * 2
		s, t = iota, iota // s, t = 7, 7
		u, v              // u, v = 8, 8
		_, w              // _, w = 9, 9
	)

	const x = iota // x = 0 （iota == 0）
	const (
		y = iota // y = 0 （iota == 0）
		z        // z = 1
	)

	println(m)             // +1.500000e+000
	println(n)             // +2.500000e+000
	println(q, r)          // 8 12
	println(s, t, u, v, w) // 7 7 8 8 9
	println(x, y, z)       // 0 0 1
}