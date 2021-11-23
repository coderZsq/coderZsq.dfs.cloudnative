package main

import "fmt"

type I interface {
	m(int)bool
}

type T string
func (t T) m(n int) bool {
	return len(t) > n
}

func main() {
	var i I = T("gopher")
	fmt.Println(i.m(5))                          // true
	fmt.Println(I.m(i, 5))                       // true
	fmt.Println(interface {m(int) bool}.m(i, 5)) // true

	// 下面这几行被执行的时候都将会产生一个恐慌。
	I(nil).m(5)
	I.m(nil, 5)
	interface {m(int) bool}.m(nil, 5)
}
