package main

import "fmt"

const s = "Go101.org" // len(s) == 9

// len(s)是一个常量表达式，但len(s[:])却不是。
var a byte = 1 << len(s) / 128
var b byte = 1 << len(s[:]) / 128

func main() {
	fmt.Println(a, b) // 4 0
}