package main

import "fmt"

func main() {
	i := 0
Next:
	if i >= 5 {
		goto Exit
	}
	// 创建一个显式代码块以缩小k的作用域。
	{
		k := i + i
		fmt.Println(k)
	}
	i++
	goto Next
Exit:
}
