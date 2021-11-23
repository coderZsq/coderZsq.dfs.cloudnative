package main

import "fmt"

func main() {
	type T struct{age int}
	mt := map[string]T{}
	mt["John"] = T{age: 29} // 整体修改是允许的
	ma := map[int][5]int{}
	ma[1] = [5]int{1: 789} // 整体修改是允许的

	// 这两个赋值编译不通过，因为部分修改一个映射
	// 元素是非法的。这看上去确实有些反直觉。
	/*
		ma[1][1] = 123      // error
		mt["John"].age = 30 // error
	*/

	// 读取映射元素的元素或者字段是没问题的。
	fmt.Println(ma[1][1])       // 789
	fmt.Println(mt["John"].age) // 29
}
