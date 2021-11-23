package main

import "fmt"

func main() {
	words := []string{
		"Go", "is", "a", "high",
		"efficient", "language.",
	}

	// fmt.Println函数的原型为：
	// func Println(a ...interface{}) (n int, err error)
	// 所以words...不能传递给此函数的调用。

	// fmt.Println(words...) // 编译不通过

	// 将[]string值转换为类型[]interface{}。
	iw := make([]interface{}, 0, len(words))
	for _, w := range words {
		iw = append(iw, w)
	}
	fmt.Println(iw...) // 编译没问题
}
