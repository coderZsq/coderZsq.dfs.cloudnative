package main

import "fmt"
import "example.com/go101/internal/22_type-embedding/case10/x.y/foo"

type B struct {
	n bool
}

func (b B) m() {
	fmt.Println("B", b.n)
}

type C struct{
	foo.A
	B
}

func main() {
	var c C
	c.m()      // B false
	foo.Bar(c) // A 0
}
