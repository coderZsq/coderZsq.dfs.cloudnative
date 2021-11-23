package main

import "fmt"

type A struct {
	x string
}
func (A) y(int) bool {
	return false
}

type B struct {
	y bool
}
func (B) x(string) {}

type C struct {
	B
}

var v2 struct {
	A
	C
}

func f2() {
	fmt.Printf("%T \n", v2.x) // string
	fmt.Printf("%T \n", v2.y) // func(int) bool
}
func main() {

}
