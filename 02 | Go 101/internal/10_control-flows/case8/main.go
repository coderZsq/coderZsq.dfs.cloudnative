package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Print(i)
		i := i // 这里声明的变量i遮挡了上面声明的i。
		// 右边的i为上面声明的循环变量i。
		i = 10 // 新声明的i被更改了。
		_ = i
	}
}
