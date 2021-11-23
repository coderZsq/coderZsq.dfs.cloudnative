package main

func main() {
	// 情况一：
	var s = make([]byte, 10000)
	y = make([]T, len(s)) // works
	copy(y, s)

	// 情况二：
	var s = make([]byte, 10000)
	y = make([]T, len(s), len(s)) // not work
	copy(y, s)

	// 情况三：
	var a = [1][]byte{s}
	y = make([]T, len(a[0])) // not work
	copy(y, a[0])

	// 情况四：
	type T struct {x []byte}
	var t = T{x: s}
	y = make([]T, len(t.x)) // not work
	copy(y, t.x)
}
