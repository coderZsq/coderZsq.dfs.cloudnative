package main

func main() {
	var a uint = 1
	var _ = map[uint]int {a : 123} // 没问题
	var _ = []int{a: 100}          // error: 下标必须为常量
	var _ = [5]int{a: 100}         // error: 下标必须为常量
}
