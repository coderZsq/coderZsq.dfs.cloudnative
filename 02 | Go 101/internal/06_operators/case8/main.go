package main

func main() {
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
}
