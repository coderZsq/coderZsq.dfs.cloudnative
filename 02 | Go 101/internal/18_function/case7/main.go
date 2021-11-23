package main

import "fmt"

func Sum(values ...int64) (sum int64) {
	sum = 0
	for _, v := range values {
		sum += v
	}
	return
}

func main() {
	a0 := Sum()
	a1 := Sum(2)
	a3 := Sum(2, 3, 5)
	// 上面三行和下面三行是等价的。
	b0 := Sum([]int64{}...) // <=> Sum(nil...)
	b1 := Sum([]int64{2}...)
	b3 := Sum([]int64{2, 3, 5}...)
	fmt.Println(a0, a1, a3) // 0 2 10
	fmt.Println(b0, b1, b3) // 0 2 10
}
