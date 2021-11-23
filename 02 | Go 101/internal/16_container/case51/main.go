package main

func main() {
	sClone := append(s[:0:0], s...)

	sClone := append([]T(nil), s...)

	// 两行make+copy实现：
	sClone := make([]T, len(s))
	copy(sClone, s)

	// 或者下面的make+append实现。
	// 对于目前的官方Go工具链v1.17来说，这种
	// 实现比上面的make+copy实现略慢一点。
	sClone := append(make([]T, 0, len(s)), s...)

	var sClone []T
	if s != nil {
		sClone = make([]T, len(s))
		copy(sClone, s)
	}
}
