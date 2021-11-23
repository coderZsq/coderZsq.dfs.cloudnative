package main

import "fmt"

func Double(n int) int {
	return n + n
}

func Apply(n int, f func(int) int) int {
	return f(n) // f的类型为"func(int) int"
}

func main() {
	fmt.Printf("%T\n", Double) // func(int) int
	// Double = nil // error: Double是不可修改的

	var f func(n int) int // 默认值为nil
	f = Double
	g := Apply
	fmt.Printf("%T\n", g) // func(int, func(int) int) int

	fmt.Println(f(9))         // 18
	fmt.Println(g(6, Double)) // 12
	fmt.Println(Apply(6, f))  // 12
}
