package solution2

func myPow(x float64, n int) float64 {
	if n < 0 { // 指数为负数, 倒数
		x = 1 / x
		n = -n
	}
	result := 1.
	for n != 0 {
		if n & 1 != 0 { // 位运算判断指数
			result *= x
		}
		x *= x
		n >>= 1 // (n = n / 2)
	}
	return result
}
