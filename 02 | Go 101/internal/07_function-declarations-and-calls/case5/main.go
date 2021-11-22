package main

func CompareLower4bits(m, n uint32) (r bool) {
	// 下面这两行等价于：return m&0xFF > n&0xff
	r = m&0xF > n&0xf
	return
}

// 此函数没有输入参数。它的结果声明列表只包含一个
// 匿名结果声明，因此它不必用()括起来。
func VersionString() string {
	return "go1.0"
}

// 此函数没有返回结果。它的所有输入参数都是匿名的。
// 它的结果声明列表为空，因此可以被省略掉。
func doNothing(string, int) {
}
