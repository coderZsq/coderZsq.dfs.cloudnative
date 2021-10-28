package main

import "fmt"

func init() {
	var (
		a, b float32 = 12.0, 3.14
		c, d int16   = 15, -6
		e    uint8   = 7
	)

	// 这些行编译没问题。
	_ = 12 + 'A' // 两个类型不确定操作数（都为数值类型）
	_ = 12 - a   // 12将被当做a的类型（float32）使用。
	_ = a * b    // 两个同类型的类型确定操作数。
	_ = c % d
	_, _ = c+int16(e), uint8(c)+e
	_, _, _, _ = a/b, c/d, -100/-9, 1.23/1.2
	_, _, _, _ = c|d, c&d, c^d, c&^d
	_, _, _, _ = d<<e, 123>>e, e>>3, 0xF<<0
	_, _, _, _ = -b, +c, ^e, ^-1

	// 这些行编译将失败。
	//_ = a % b   // error: a和b都不是整数
	//_ = a | b   // error: a和b都不是整数
	//_ = c + e   // error: c和e的类型不匹配
	//_ = b >> 5  // error: b不是一个整数
	//_ = c >> -5 // error: -5不是一个无符号整数

	_ = e << uint(c) // 编译没问题
	_ = e << c       // 从Go 1.13开始，此行才能编译过
	//_ = e << -c      // 从Go 1.13开始，此行才能编译过。
	// 将在运行时刻造成恐慌。
	//_ = e << -1      // error: 右操作数不能为负（常数）
}

func init() {
	// 结果为非常量
	var a, b uint8 = 255, 1
	var c = a + b // c==0。a+b是一个非常量表达式，
	// 结果中溢出的高位比特将被截断舍弃。
	var d = a << b // d == 254。同样，结果中溢出的
	// 高位比特将被截断舍弃。

	// 结果为类型不确定常量，允许溢出其默认类型。
	const X = 0x1FFFFFFFF * 0x1FFFFFFFF // 没问题，尽管X溢出
	const R = 'a' + 0x7FFFFFFF          // 没问题，尽管R溢出

	// 运算结果或者转换结果为类型确定常量
	//var e = X // error: X溢出int。
	//var h = R // error: R溢出rune。
	//const Y = 128 - int8(1)  // error: 128溢出int8。
	//const Z = uint8(255) + 1 // error: 256溢出uint8。
	_, _ = c, d
}

func init() {
	// 三个类型不确定常量。它们的默认类型
	// 分别为：int、rune和complex64.
	const X, Y, Z = 2, 'A', 3i

	var a, b int = X, Y // 两个类型确定值

	// 变量d的类型被推断为Y的默认类型：rune（亦即int32）。
	d := X + Y
	// 变量e的类型被推断为a的类型：int。
	e := Y - a
	// 变量f的类型和a及b的类型一样：int。
	f := a * b
	// 变量g的类型被推断为Z的默认类型：complex64。
	g := Z * Y

	// 2 65 (+0.000000e+000+3.000000e+000i)
	println(X, Y, Z)
	// 67 63 130 (+0.000000e+000+1.950000e+002i)
	println(d, e, f, g)
}

func init() {
	const N = 2
	// A == 12，它是一个默认类型为int的类型不确定值。
	const A = 3.0 << N
	// B == 12，它是一个类型为int8的类型确定值。
	const B = int8(3.0) << N

	var m = uint(32)
	// 下面的三行是相互等价的。
	var x int64 = 1 << m  // 1的类型将被设想为int64，而非int
	var y = int64(1 << m) // 同上
	var z = int64(1) << m // 同上

	// 下面这行编译不通过。
	/*
	   var _ = 1.23 << m // error: 浮点数不能被移位
	*/

	_, _, _ = x, y, z
}

func init() {
	const n = uint(2)
	var m = uint(2)

	// 这两行编译没问题。
	var _ float64 = 1 << n
	var _ = float64(1 << n)

	// 这两行编译失败。
	//var _ float64 = 1 << m  // error
	//var _ = float64(1 << m) // error

	_ = m
}

func init() {
	//var _ = float64(1) << m
	//var _ = 1.0 << m // error: shift of type float64
}

const n = uint(8)

var m = uint(8)

func init() {
	println(a, b) // 2 0
}

var a byte = 1 << n / 128
var b byte = 1 << m / 128

func init() {
	var a = byte(int(1) << n / 128)
	var b = byte(1) << m / 128
	println(a, b)
}

func init() {
	println(5/3, 5%3)     // 1 2
	println(5/-3, 5%-3)   // -1 2
	println(-5/3, -5%3)   // -1 -2
	println(-5/-3, -5%-3) // 1 -2

	println(5.0 / 3.0)           // 1.666667
	println((1 - 1i) / (1 + 1i)) // -1i

	var a, b = 1.0, 0.0
	println(a/b, b/b) // +Inf NaN

	//_ = int(a)/int(b) // 编译没问题，但在运行时刻将造成恐慌。

	// 这两行编译不通过。
	//println(1.0/0.0) // error: 除数为0
	//println(0.0/0.0) // error: 除数为0
}

func init() {
	var a, b int8 = 3, 5
	a += b
	println(a) // 8
	a *= a
	println(a) // 64
	a /= b
	println(a) // 12
	a %= b
	println(a) // 2
	b <<= uint(a)
	println(b) // 20
	b >>= uint(a)
	println(b) // 5
}

func init() {
	a, b, c := 12, 1.2, 1+2i
	a++ // ok. <=> a += 1 <=> a = a + 1
	b-- // ok. <=> b -= 1 <=> b = b - 1
	c++ // ok.

	// 下面这些行编译不通过。
	/*
		_ = a++
		_ = b--
		_ = c++
		++a
		--b
		++c
	*/
}

func init() {
	println("Go" + "lang") // Golang
	var a = "Go"
	a += "lang"
	println(a) // Golang
}

func init() {
	var x = 1.2 + 3/2
	fmt.Println(x)
}

func init() {
	const x = 3 / 2 * 0.1
	const y = 0.1 * 3 / 2

	println(x) // +1.000000e-001
	println(y) // +1.500000e-001
}

func main() {}
