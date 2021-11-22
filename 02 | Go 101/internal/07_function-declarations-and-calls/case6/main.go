package main

func SquaresOfSumAndDiff(a int64, b int64) (int64, int64) {
	return (a+b) * (a+b), (a-b) * (a-b)
}

func CompareLower4bits(m, n uint32) (r bool) {
	r = m&0xF > n&0xf
	return
}

// 使用一个函数调用的返回结果来初始化一个包级变量。
var v = VersionString()

func main() {
	println(v) // v1.0
	x, y := SquaresOfSumAndDiff(3, 6)
	println(x, y) // 81 9
	b := CompareLower4bits(uint32(x), uint32(y))
	println(b) // false
	// "Go"的类型被推断为string；1的类型被推断为int32。
	doNothing("Go", 1)
}

func VersionString() string {
	return "v1.0"
}

func doNothing(string, int32) {
}