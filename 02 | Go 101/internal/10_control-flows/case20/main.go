package main

import "fmt"

func main() {
	i := 0
Next:
	if i >= 5 {
		// error: goto Exit jumps over declaration of k
		goto Exit
	}

	k := i + i
	fmt.Println(k)
	i++
	goto Next
Exit: // 此标签声明在k的作用域内，但
	// 它的使用在k的作用域之外。
}
