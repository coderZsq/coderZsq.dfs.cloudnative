package main

func main() {
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
