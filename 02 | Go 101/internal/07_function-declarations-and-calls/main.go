package main

func SquaresOfSumAndDiff(a int64, b int64) (s int64, d int64) {
	x, y := a + b, a - b
	s = x * x
	d = y * y
	return // <=> return s, d
}

func SquaresOfSumAndDiff2(a int64, b int64) (int64, int64) {
	return (a+b) * (a+b), (a-b) * (a-b)
}

func f() (x int, y bool) {
	println(x, y) // 0 false
	return
}

func SquaresOfSumAndDiff3(a, b int64) (s, d int64) {
	return (a+b) * (a+b), (a-b) * (a-b)
	// 上面这行等价于下面这行：
	// s = (a+b) * (a+b); d = (a-b) * (a-b); return
}

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
func doNothing(string, int32) {
}

// 使用一个函数调用的返回结果来初始化一个包级变量。
var v = VersionString()

func init() {
	println(v) // v1.0
	x, y := SquaresOfSumAndDiff2(3, 6)
	println(x, y) // 81 9
	b := CompareLower4bits(uint32(x), uint32(y))
	println(b) // false
	// "Go"的类型被推断为string；1的类型被推断为int32。
	doNothing("Go", 1)
}

func init() {
	// 这个匿名函数没有输入参数，但有两个返回结果。
	x, y := func() (int, int) {
		println("This fucntion has no parameters.")
		return 3, 4
	}() // 一对小括号表示立即调用此函数。不需传递实参。

	// 下面这些匿名函数没有返回结果。

	func(a, b int) {
		println("a*a + b*b =", a*a + b*b) // a*a + b*b = 25
	}(x, y) // 立即调用并传递两个实参。

	func(x int) {
		// 形参x遮挡了外层声明的变量x。
		println("x*x + y*y =", x*x + y*y) // x*x + y*y = 32
	}(y) // 将实参y传递给形参x。

	func() {
		println("x*x + y*y =", x*x + y*y) // x*x + y*y = 25
	}() // 不需传递实参。
}

func init() {
	// c是一个类型不确定复数常量。
	const c = complex(1.6, 3.3)

	// 函数调用real(c)和imag(c)的结果都是类型
	// 不确定浮点数值。在下面这句赋值中，它们都
	// 被推断为float32类型的值。
	var a, b float32 = real(c), imag(c)

	// 变量d的类型被推断为内置类型complex64。
	// 函数调用real(d)和imag(d)的结果都是
	// 类型为float32的类型确定值。
	var d = complex(a, b)

	// 变量e的类型被推断为内置类型complex128。
	// 函数调用real(e)和imag(e)的结果都是
	// 类型为float64的类型确定值。
	var e = c

	_, _ = d, e
}

func main() {}