package main

func main() {
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
