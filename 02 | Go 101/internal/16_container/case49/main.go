package main

import "fmt"

var a [5]int
var p *[7]string

// N和M都是类型为int的类型确定值。
const N = len(a)
const M = cap(p)

func main() {
	fmt.Println(N) // 5
	fmt.Println(M) // 7
}
