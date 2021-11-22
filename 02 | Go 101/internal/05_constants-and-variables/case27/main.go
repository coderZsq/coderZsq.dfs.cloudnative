package main

// 三个类型不确定常量。
const n = 1 << 64          // 默认类型为int
const r = 'a' + 0x7FFFFFFF // 默认类型为rune
const x = 2e+308           // 默认类型为float64

func main() {
	_ = n >> 2
	_ = r - 0x7FFFFFFF
	_ = x / 2
}
