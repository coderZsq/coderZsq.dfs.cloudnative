package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch n := rand.Intn(100) % 5; n {
	case 0, 1, 2, 3, 4:
		fmt.Println("n =", n)
		// 此整个if代码块为当前分支中的最后一条语句
		if true {
			//fallthrough // error: 不是当前分支中的最后一条语句
		}
	case 5, 6, 7, 8:
		n := 99
		//fallthrough // error: 不是当前分支中的最后一条语句
		_ = n
	default:
		fmt.Println(n)
		//fallthrough // error: 不能出现在最后一个分支中
	}
}
