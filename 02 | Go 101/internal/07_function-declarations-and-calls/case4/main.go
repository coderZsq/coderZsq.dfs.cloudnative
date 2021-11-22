package main

func SquaresOfSumAndDiff(a, b int64) (s, d int64) {
	return (a+b) * (a+b), (a-b) * (a-b)
	// 上面这行等价于下面这行：
	// s = (a+b) * (a+b); d = (a-b) * (a-b); return
}
