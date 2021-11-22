package main

import "fmt"

func main() {
	for i := 0; ; i++ { // 等价于：for i := 0; true; i++ {
		if i >= 10 {
			break
		}
		fmt.Println(i)
	}

	// 下面这几个循环是等价的。
	for ; true; {
	}
	for true {
	}
	for ; ; {
	}
	for {
	}
}
