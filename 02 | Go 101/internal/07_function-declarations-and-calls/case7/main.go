package main

func main() {
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
