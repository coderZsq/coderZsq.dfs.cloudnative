package main

import (
	"example.com/go101/internal/22_type-embedding/case4/foo"
	"fmt"
)

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
	//var c C
	//c.m()      // B false
	//foo.Bar(c) // A 0
}

//// 注意：这些声明不是合法的Go语法。这里这样表示只是为了
//// 解释目的。它们有助于解释提升方法值是如何被估值的。
//func (s Singer) PrintName = s.Person.PrintName
//func (s *Singer) PrintName = s.Person.PrintName
//func (s *Singer) SetAge = s.Person.SetAge