package main

func main() {
	const N = 123
	var x int
	var y, z float32

	//N = 789 // error: N是一个不可变量
	y = N   // ok: N被隐式转换为类型float32
	//x = y   // error: 类型不匹配
	x = N   // ok: N被隐式转换为类型int
	//y = x   // error: 类型不匹配
	z = y   // ok
	_ = y   // ok

	z, y = y, z               // ok
	_, y = y, z               // ok
	z, _ = y, z               // ok
	_, _ = y, z               // ok
	x, y = 69, 1.23           // ok
	//x, y = y, x               // error: 类型不匹配
	x, y = int(y), float32(x) // ok
}
