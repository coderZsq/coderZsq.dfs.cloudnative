package main

import "fmt"

func double(x *int) {
	*x += *x
	x = nil // 此行仅为讲解目的
}

func main() {
	var a = 3
	double(&a)
	fmt.Println(a) // 6
	p := &a
	double(p)
	fmt.Println(a, p == nil) // 12 false
}