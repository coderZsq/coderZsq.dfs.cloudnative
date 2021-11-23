package main

import "fmt"

func main() {
	a := int64(5)
	p := &a

	// 下面这两行编译不通过。
	/*
		p++
		p = (&a) + 8
	*/

	*p++
	fmt.Println(*p, a)   // 6 6
	fmt.Println(p == &a) // true

	*&a++
	*&*&a++
	**&p++
	*&*p++
	fmt.Println(*p, a) // 10 10
}

