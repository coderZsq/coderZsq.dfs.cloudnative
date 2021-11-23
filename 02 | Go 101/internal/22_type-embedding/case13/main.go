package main

import "fmt"

type X int

func (x X) M1() {
	fmt.Println(x)
}

func (x *X) M2() {
	fmt.Println(*x)
}

type T struct { X }

type S struct { *T }

func main() {
	var t = &T{X: 1}
	var s = S{T: t}
	var f = s.M1 // <=> (*s.T).X.M1
	var g = s.M2 // <=> (&(*s.T).X).M2
	s.X = 2
	f() // 1
	g() // 2
	s.T = &T{X: 3}
	f() // 1
	g() // 2
}
