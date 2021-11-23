package main

import "fmt"

func main() {
	var a, b, c interface{} = "abc", 123, "a"+"b"+"c"
	fmt.Println(a == b) // 第二步的情形。输出"false"。
	fmt.Println(a == c) // 第三步的情形。输出"true"。

	var x *int = nil
	var y *bool = nil
	var ix, iy interface{} = x, y
	var i interface{} = nil
	fmt.Println(ix == iy) // 第二步的情形。输出"false"。
	fmt.Println(ix == i)  // 第一步的情形。输出"false"。
	fmt.Println(iy == i)  // 第一步的情形。输出"false"。

	var s []int = nil // []int为一个不可比较类型。
	i = s
	fmt.Println(i == nil) // 第一步的情形。输出"false"。
	fmt.Println(i == i)   // 第三步的情形。将产生一个恐慌。
}
