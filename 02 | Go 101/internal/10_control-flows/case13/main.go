package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(100) % 5; n {
	case 0, 1, 2, 3, 4:
		fmt.Println("n =", n)
		fallthrough // 跳到下个代码块
	case 5, 6, 7, 8:
		// 一个新声明的n，它只在当前分支代码快内可见。
		n := 99
		fmt.Println("n =", n) // 99
		fallthrough
	default:
		// 下一行中的n和第一个分支中的n是同一个变量。
		// 它们均为switch表达式"n"。
		fmt.Println("n =", n)
	}
}
