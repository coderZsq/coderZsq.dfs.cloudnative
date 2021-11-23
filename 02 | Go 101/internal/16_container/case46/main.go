package main

import "fmt"

func main() {
	pa := &[5]int{2, 3, 5, 7, 11}
	s := pa[1:3]
	fmt.Println(s) // [3 5]
	pa = nil
	s = pa[0:0] // panic
	// 如果下一行能被执行到，则它也会产生恐慌。
	_ = (*[0]byte)(nil)[:]
}
