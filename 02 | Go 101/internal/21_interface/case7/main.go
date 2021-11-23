package main

import "fmt"

func main() {
	var i interface{}
	i = []int{1, 2, 3}
	fmt.Println(i) // [1 2 3]
	i = map[string]int{"Go": 2012}
	fmt.Println(i) // map[Go:2012]
	i = true
	fmt.Println(i) // true
	i = 1
	fmt.Println(i) // 1
	i = "abc"
	fmt.Println(i) // abc

	// 将接口值i中包裹的值清除掉。
	i = nil
	fmt.Println(i) //
}
