package main

func main() {
	println( 5/3,   5%3)  // 1 2
	println( 5/-3,  5%-3) // -1 2
	println(-5/3,  -5%3)  // -1 -2
	println(-5/-3, -5%-3) // 1 -2

	println(5.0 / 3.0)     // 1.666667
	println((1-1i)/(1+1i)) // -1i

	var a, b = 1.0, 0.0
	println(a/b, b/b) // +Inf NaN

	_ = int(a)/int(b) // 编译没问题，但在运行时刻将造成恐慌。

	// 这两行编译不通过。
	println(1.0/0.0) // error: 除数为0
	println(0.0/0.0) // error: 除数为0
}