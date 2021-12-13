package solution2

func myPow(x float64, n int) float64 {
	if n == 0 { // 指数为 0
		return 1
	}
	if n < 0 { // 指数为负数, 返回倒数
		return 1 / myPow(x, -n) // 此时的 -n 为正数
	}
	if n%2 == 0 { // 指数为偶数
		return myPow(x*x, n/2) // 分治
	}
	if n%2 != 0 { // 指数为奇数
		return myPow(x, n-1) * x // 将其先变为偶数
	}
	return -1
}
