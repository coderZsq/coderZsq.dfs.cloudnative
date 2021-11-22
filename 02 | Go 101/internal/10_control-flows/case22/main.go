package main

import "fmt"

func main() {
	var k int // 将变量k的声明移到此处。
	i := 0
Next:
	if i >= 5 {
		goto Exit
	}

	k = i + i
	fmt.Println(k)
	i++
	goto Next
Exit:
}
